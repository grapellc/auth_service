package app

import (
	"context"
	"crypto/rand"
	"fmt"
	"io"
	"regexp"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
	"github.com/sirupsen/logrus"
	"github.com/your-moon/grape-auth-service/internal/auth/errs"
	"github.com/your-moon/grape-auth-service/internal/auth/ports"
	"github.com/your-moon/grape-shared/entities"
	"golang.org/x/crypto/bcrypt"
)

const otpTTL = 5 * time.Minute

const (
	registrationOTPLimit   = 10
	forgotPasswordOTPLimit = 8
	maxOTPVerifyAttempts   = 5
	maxLoginAttempts       = 5
)

var (
	emailRegex = regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	phoneRegex = regexp.MustCompile(`^[0-9]{8,15}$`)
)

type Service struct {
	userRepo         ports.UserRepository
	authLogRepo      ports.AuthLogRepository
	refreshTokenRepo ports.RefreshTokenRepository
	otpSender        ports.OTPSender
	redis            redis.Cmdable
	tokenService     ports.TokenService
	codeGenerator    func() string
	clock            func() time.Time
}

func NewService(
	userRepo ports.UserRepository,
	authLogRepo ports.AuthLogRepository,
	refreshTokenRepo ports.RefreshTokenRepository,
	otpSender ports.OTPSender,
	redis redis.Cmdable,
	tokenService ports.TokenService,
) *Service {
	return &Service{
		userRepo:         userRepo,
		authLogRepo:      authLogRepo,
		refreshTokenRepo: refreshTokenRepo,
		otpSender:        otpSender,
		redis:            redis,
		tokenService:     tokenService,
		clock:            time.Now,
		codeGenerator: func() string {
			var table = [...]byte{'1', '2', '3', '4', '5', '6', '7', '8', '9', '0'}
			b := make([]byte, 6)
			n, err := io.ReadAtLeast(rand.Reader, b, 6)
			if n != 6 || err != nil {
				return fmt.Sprintf("%06d", time.Now().UnixNano()%1000000)
			}
			for i := 0; i < len(b); i++ {
				b[i] = table[int(b[i])%len(table)]
			}
			return string(b)
		},
	}
}

func (s *Service) SetCodeGenerator(generator func() string) {
	s.codeGenerator = generator
}

func (s *Service) SetClock(clock func() time.Time) {
	s.clock = clock
}

func (s *Service) LogAuth(ctx context.Context, log *entities.AuthLogUUID) {
	if err := s.authLogRepo.Create(ctx, log); err != nil {
		logrus.Errorf("Failed to create auth log: %v", err)
	}
}

func (s *Service) normalizeIdentifier(identifier string) string {
	return strings.ToLower(strings.TrimSpace(identifier))
}

func (s *Service) isEmail(identifier string) bool {
	return emailRegex.MatchString(identifier)
}

func (s *Service) isPhone(identifier string) bool {
	return phoneRegex.MatchString(identifier)
}

func (s *Service) RequestOTP(ctx context.Context, identifier string) error {
	identifier = s.normalizeIdentifier(identifier)
	if !s.isEmail(identifier) && !s.isPhone(identifier) {
		return errs.ErrInvalidIdentifier
	}
	rateLimitKey := fmt.Sprintf("otp_rate:default:%s", identifier)
	if err := s.checkAndIncrementRateLimit(ctx, rateLimitKey, 10); err != nil {
		return err
	}
	return s.sendOTP(ctx, identifier)
}

func (s *Service) sendOTP(ctx context.Context, identifier string) error {
	code := s.codeGenerator()
	logrus.WithFields(logrus.Fields{"identifier": identifier, "code": code, "ttl": otpTTL}).Info("Generated OTP")
	key := fmt.Sprintf("otp:%s", identifier)
	if err := s.redis.Set(ctx, key, code, otpTTL).Err(); err != nil {
		logrus.Errorf("Failed to store OTP in Redis: %v", err)
		return err
	}
	attemptKey := fmt.Sprintf("otp_attempts:%s", identifier)
	s.redis.Del(ctx, attemptKey)
	err := s.otpSender.Send(ctx, identifier, code)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"identifier": identifier,
			"error":      err,
		}).Error("OTP sender is gone or down")
	}
	return err
}

func (s *Service) VerifyOTP(ctx context.Context, identifier, code string, consume bool) error {
	identifier = s.normalizeIdentifier(identifier)
	key := fmt.Sprintf("otp:%s", identifier)
	attemptKey := fmt.Sprintf("otp_attempts:%s", identifier)
	attempts, _ := s.redis.Get(ctx, attemptKey).Int()
	if attempts >= maxOTPVerifyAttempts {
		return errs.ErrTooManyAttempts
	}
	storedCode, err := s.redis.Get(ctx, key).Result()
	if err == redis.Nil {
		return errs.ErrOTPExpired
	}
	if err != nil {
		return err
	}
	if storedCode != code {
		s.redis.Incr(ctx, attemptKey)
		s.redis.Expire(ctx, attemptKey, otpTTL)
		return errs.ErrInvalidOTP
	}
	if consume {
		s.redis.Del(ctx, key)
		s.redis.Del(ctx, attemptKey)
	}
	return nil
}

func (s *Service) SignUp(ctx context.Context, identifier, code, password string) (*entities.UserUUID, string, string, error) {
	identifier = s.normalizeIdentifier(identifier)
	if err := s.VerifyOTP(ctx, identifier, code, true); err != nil {
		return nil, "", "", err
	}
	existingUser, _ := s.userRepo.FindByPhoneOrEmail(identifier)
	if existingUser != nil {
		return nil, "", "", errs.ErrUserAlreadyExists
	}
	isEmail := s.isEmail(identifier)
	var email string
	var phone *string
	if isEmail {
		email = identifier
		phone = nil
	} else {
		phone = &identifier
		email = identifier + "@placeholder.com"
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", "", err
	}
	pwHash := string(hashedPassword)
	newUser := &entities.UserUUID{
		PhoneNumber: phone, Email: email, PasswordHash: &pwHash,
		IsPhoneVerified: !isEmail, IsEmailVerified: isEmail, Role: "user",
	}
	if err := s.userRepo.Create(newUser); err != nil {
		return nil, "", "", err
	}
	return s.generateTokenPair(newUser)
}

func (s *Service) VerifyRegistration(ctx context.Context, identifier, code string) (*entities.UserUUID, string, string, error) {
	identifier = s.normalizeIdentifier(identifier)
	if err := s.VerifyOTP(ctx, identifier, code, true); err != nil {
		return nil, "", "", err
	}
	existingUser, _ := s.userRepo.FindByPhoneOrEmail(identifier)
	if existingUser != nil {
		if existingUser.PasswordHash != nil {
			return nil, "", "", errs.ErrUserAlreadyExists
		}
		return s.generateTokenPair(existingUser)
	}
	isEmail := s.isEmail(identifier)
	var email string
	var phone *string
	if isEmail {
		email = identifier
		phone = nil
	} else {
		phone = &identifier
		email = identifier + "@placeholder.com"
	}
	newUser := &entities.UserUUID{
		PhoneNumber: phone, Email: email, PasswordHash: nil,
		IsPhoneVerified: !isEmail, IsEmailVerified: isEmail, Role: "user",
	}
	if err := s.userRepo.Create(newUser); err != nil {
		return nil, "", "", err
	}
	return s.generateTokenPair(newUser)
}

func (s *Service) SetPassword(ctx context.Context, userID string, password string) error {
	id, err := uuid.Parse(userID)
	if err != nil {
		return errs.ErrUserNotFound
	}
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return errs.ErrUserNotFound
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	pwHash := string(hashedPassword)
	user.PasswordHash = &pwHash
	if user.PhoneNumber != nil {
		user.IsPhoneVerified = true
	}
	if len(user.Email) <= 16 || user.Email[len(user.Email)-16:] != "@placeholder.com" {
		user.IsEmailVerified = true
	}
	return s.userRepo.Update(user)
}

func (s *Service) LoginPassword(ctx context.Context, identifier, password string) (*entities.UserUUID, string, string, error) {
	identifier = s.normalizeIdentifier(identifier)
	loginAttemptKey := fmt.Sprintf("login_attempts:%s", identifier)
	attempts, _ := s.redis.Get(ctx, loginAttemptKey).Int()
	if attempts >= maxLoginAttempts {
		return nil, "", "", errs.ErrTooManyAttempts
	}
	user, err := s.userRepo.FindByPhoneOrEmail(identifier)
	if err != nil {
		s.redis.Incr(ctx, loginAttemptKey)
		s.redis.Expire(ctx, loginAttemptKey, 15*time.Minute)
		return nil, "", "", errs.ErrInvalidCredentials
	}
	if user.PasswordHash == nil {
		return user, "", "", errs.ErrPasswordNotSet
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*user.PasswordHash), []byte(password)); err != nil {
		s.redis.Incr(ctx, loginAttemptKey)
		s.redis.Expire(ctx, loginAttemptKey, 15*time.Minute)
		return user, "", "", errs.ErrInvalidCredentials
	}
	s.redis.Del(ctx, loginAttemptKey)
	return s.generateTokenPair(user)
}

func (s *Service) LoginOTP(ctx context.Context, phone, code string) (*entities.UserUUID, string, string, error) {
	if err := s.VerifyOTP(ctx, phone, code, true); err != nil {
		return nil, "", "", err
	}
	user, err := s.userRepo.FindByPhone(phone)
	if err != nil {
		return nil, "", "", errs.ErrUserNotFound
	}
	return s.generateTokenPair(user)
}

func (s *Service) ResetPassword(ctx context.Context, identifier, code, newPassword string) error {
	identifier = s.normalizeIdentifier(identifier)
	if err := s.VerifyOTP(ctx, identifier, code, true); err != nil {
		return err
	}
	user, err := s.userRepo.FindByPhoneOrEmail(identifier)
	if err != nil {
		return errs.ErrUserNotFound
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	pwHash := string(hashedPassword)
	user.PasswordHash = &pwHash
	if s.isPhone(identifier) {
		user.IsPhoneVerified = true
	} else if s.isEmail(identifier) {
		user.IsEmailVerified = true
	}
	return s.userRepo.Update(user)
}

func (s *Service) RequestPhoneChange(ctx context.Context, userID string, newPhone string) error {
	uid, _ := uuid.Parse(userID)
	existingUser, _ := s.userRepo.FindByPhone(newPhone)
	if existingUser != nil && existingUser.ID != uid && existingUser.IsPhoneVerified {
		return errs.ErrPhoneAlreadyExists
	}
	return s.RequestOTP(ctx, newPhone)
}

func (s *Service) ConfirmPhoneChange(ctx context.Context, userID string, newPhone, code string) error {
	if err := s.VerifyOTP(ctx, newPhone, code, true); err != nil {
		return err
	}
	uid, err := uuid.Parse(userID)
	if err != nil {
		return errs.ErrUserNotFound
	}
	existingUser, _ := s.userRepo.FindByPhone(newPhone)
	if existingUser != nil && existingUser.ID != uid && existingUser.IsPhoneVerified {
		return errs.ErrPhoneAlreadyExists
	}
	user, err := s.userRepo.GetByID(uid)
	if err != nil {
		return errs.ErrUserNotFound
	}
	user.PhoneNumber = &newPhone
	user.IsPhoneVerified = true
	return s.userRepo.Update(user)
}

func (s *Service) RequestEmailChange(ctx context.Context, userID string, newEmail string) error {
	newEmail = s.normalizeIdentifier(newEmail)
	if !s.isEmail(newEmail) {
		return errs.ErrInvalidEmail
	}
	uid, _ := uuid.Parse(userID)
	existingUser, _ := s.userRepo.FindByEmail(newEmail)
	if existingUser != nil && existingUser.ID != uid && existingUser.IsEmailVerified {
		return errs.ErrEmailAlreadyExists
	}
	return s.RequestOTP(ctx, newEmail)
}

func (s *Service) ConfirmEmailChange(ctx context.Context, userID string, newEmail, code string) error {
	newEmail = s.normalizeIdentifier(newEmail)
	if err := s.VerifyOTP(ctx, newEmail, code, true); err != nil {
		return err
	}
	uid, err := uuid.Parse(userID)
	if err != nil {
		return errs.ErrUserNotFound
	}
	existingUser, _ := s.userRepo.FindByEmail(newEmail)
	if existingUser != nil && existingUser.ID != uid && existingUser.IsEmailVerified {
		return errs.ErrEmailAlreadyExists
	}
	user, err := s.userRepo.GetByID(uid)
	if err != nil {
		return errs.ErrUserNotFound
	}
	user.Email = newEmail
	user.IsEmailVerified = true
	return s.userRepo.Update(user)
}

func (s *Service) checkAndIncrementRateLimit(ctx context.Context, key string, limit int) error {
	count, err := s.redis.Get(ctx, key).Int()
	if err != nil && err != redis.Nil {
		return err
	}
	if count >= limit {
		return errs.ErrTooManyAttempts
	}
	if err := s.redis.Incr(ctx, key).Err(); err != nil {
		return err
	}
	if count == 0 {
		now := s.clock()
		midnight := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
		ttl := midnight.Sub(now)
		s.redis.Expire(ctx, key, ttl)
	}
	return nil
}

func (s *Service) RequestRegistrationOTP(ctx context.Context, identifier string) error {
	identifier = s.normalizeIdentifier(identifier)
	if !s.isEmail(identifier) && !s.isPhone(identifier) {
		return errs.ErrInvalidIdentifier
	}
	existingUser, _ := s.userRepo.FindByPhoneOrEmail(identifier)
	if existingUser != nil && existingUser.PasswordHash != nil {
		logrus.Infof("Registration requested for existing user: %s", identifier)
		return nil
	}
	rateLimitKey := fmt.Sprintf("otp_rate:registration:%s", identifier)
	if err := s.checkAndIncrementRateLimit(ctx, rateLimitKey, registrationOTPLimit); err != nil {
		return err
	}
	return s.sendOTP(ctx, identifier)
}

func (s *Service) RequestForgotPasswordOTP(ctx context.Context, identifier string) error {
	identifier = s.normalizeIdentifier(identifier)
	if !s.isPhone(identifier) && !s.isEmail(identifier) {
		return errs.ErrInvalidIdentifier
	}
	rateLimitKey := fmt.Sprintf("otp_rate:forgot_password:%s", identifier)
	if err := s.checkAndIncrementRateLimit(ctx, rateLimitKey, forgotPasswordOTPLimit); err != nil {
		return err
	}
	return s.sendOTP(ctx, identifier)
}

func (s *Service) DeleteAccount(ctx context.Context, userID string) error {
	id, err := uuid.Parse(userID)
	if err != nil {
		return errs.ErrUserNotFound
	}
	user, err := s.userRepo.GetByID(id)
	if err != nil {
		return errs.ErrUserNotFound
	}
	return s.userRepo.Delete(user)
}

func (s *Service) generateTokenPair(user *entities.UserUUID) (*entities.UserUUID, string, string, error) {
	accessToken, err := s.tokenService.GenerateAccessToken(user.ID.String(), user.Role)
	if err != nil {
		return nil, "", "", err
	}
	refreshToken, err := s.tokenService.GenerateRefreshToken()
	if err != nil {
		return nil, "", "", err
	}
	rt := &entities.RefreshTokenUUID{
		ID:        uuid.New(),
		UserID:    user.ID,
		Token:     refreshToken,
		ExpiresAt: time.Now().Add(7 * 24 * time.Hour),
		Revoked:   false,
	}
	if err := s.refreshTokenRepo.Create(rt); err != nil {
		return nil, "", "", err
	}
	return user, accessToken, refreshToken, nil
}

func (s *Service) RefreshToken(ctx context.Context, token string) (string, string, error) {
	rt, err := s.refreshTokenRepo.GetByToken(token)
	if err != nil {
		return "", "", errs.ErrInvalidRefreshToken
	}
	user, err := s.userRepo.GetByID(rt.UserID)
	if err != nil {
		return "", "", errs.ErrUserNotFound
	}
	s.refreshTokenRepo.Revoke(token)
	_, newAccess, newRefresh, err := s.generateTokenPair(user)
	return newAccess, newRefresh, err
}

func (s *Service) Logout(ctx context.Context, token string) error {
	return s.refreshTokenRepo.Revoke(token)
}

func (s *Service) GetUserByID(ctx context.Context, userID string) (*entities.UserUUID, error) {
	if userID == "" {
		return nil, errs.ErrUserNotFound
	}
	id, err := uuid.Parse(userID)
	if err != nil {
		return nil, errs.ErrUserNotFound
	}
	return s.userRepo.GetByID(id)
}

// Ensure Service implements ports.AuthService.
var _ ports.AuthService = (*Service)(nil)
