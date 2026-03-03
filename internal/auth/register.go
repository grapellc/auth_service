package auth

import (
	"github.com/redis/go-redis/v9"
	"github.com/your-moon/grape-auth-service/internal/auth/adapters"
	"github.com/your-moon/grape-auth-service/internal/auth/app"
	"github.com/your-moon/grape-auth-service/internal/auth/ports"
	"github.com/your-moon/grape-auth-service/internal/otp"
	"gorm.io/gorm"
)

// NewAuthService builds the in-process auth service for the auth-service binary.
func NewAuthService(db *gorm.DB, redisClient redis.Cmdable, otpSender otp.Sender) ports.AuthService {
	userRepo := adapters.NewUserRepositoryAdapter(db)
	authLogRepo := adapters.NewAuthLogRepositoryAdapter(db)
	refreshTokenRepo := adapters.NewRefreshTokenRepositoryAdapter(db)
	otpAdapter := adapters.NewOTPSenderAdapter(otpSender)
	tokenService := adapters.NewJWTService()
	return app.NewService(userRepo, authLogRepo, refreshTokenRepo, otpAdapter, redisClient, tokenService)
}
