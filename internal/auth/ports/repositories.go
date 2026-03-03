package ports

import (
	"context"

	"github.com/google/uuid"
	"github.com/your-moon/grape-shared/entities"
)

// UserRepository defines user persistence for the auth context (UUID schema).
type UserRepository interface {
	GetByID(id uuid.UUID) (*entities.UserUUID, error)
	FindByPhone(phone string) (*entities.UserUUID, error)
	FindByEmail(email string) (*entities.UserUUID, error)
	FindByPhoneOrEmail(identifier string) (*entities.UserUUID, error)
	Create(user *entities.UserUUID) error
	Update(user *entities.UserUUID) error
	Delete(user *entities.UserUUID) error
}

// AuthLogRepository persists auth audit logs (UUID schema).
type AuthLogRepository interface {
	Create(ctx context.Context, log *entities.AuthLogUUID) error
}

// RefreshTokenRepository persists refresh tokens (UUID schema).
type RefreshTokenRepository interface {
	Create(token *entities.RefreshTokenUUID) error
	GetByToken(token string) (*entities.RefreshTokenUUID, error)
	Revoke(token string) error
	RevokeAllForUser(userID uuid.UUID) error
}

// OTPSender sends OTP codes (SMS/email).
type OTPSender interface {
	Send(ctx context.Context, to string, code string) error
}

// TokenService generates and validates JWT tokens. userID is UUID string.
type TokenService interface {
	GenerateAccessToken(userID string, role string) (string, error)
	GenerateRefreshToken() (string, error)
}
