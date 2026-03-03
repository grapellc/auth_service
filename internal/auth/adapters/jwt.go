package adapters

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"github.com/your-moon/grape-auth-service/internal/auth/ports"
)

type claims struct {
	UserID string `json:"user_id"` // UUID string
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey      []byte
	expireDuration time.Duration
}

func NewJWTService() ports.TokenService {
	secret := viper.GetString("jwt.secret")
	if secret == "" {
		logrus.Fatal("jwt.secret is not set in configuration")
	}
	return &jwtService{
		secretKey:      []byte(secret),
		expireDuration: time.Minute * 15,
	}
}

func (s *jwtService) GenerateAccessToken(userID string, role string) (string, error) {
	c := &claims{
		UserID: userID,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.expireDuration)),
			Issuer:    "grape-backend",
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	return token.SignedString(s.secretKey)
}

func (s *jwtService) GenerateRefreshToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return hex.EncodeToString(b), nil
}
