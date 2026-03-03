package adapters

import (
	"context"

	"github.com/your-moon/grape-auth-service/internal/auth/ports"
	authrepo "github.com/your-moon/grape-auth-service/internal/auth/repository"
	"github.com/your-moon/grape-auth-service/internal/otp"
	"gorm.io/gorm"
)

func NewUserRepositoryAdapter(db *gorm.DB) ports.UserRepository {
	return authrepo.NewUserRepository(db)
}

func NewAuthLogRepositoryAdapter(db *gorm.DB) ports.AuthLogRepository {
	return authrepo.NewAuthLogRepository(db)
}

func NewRefreshTokenRepositoryAdapter(db *gorm.DB) ports.RefreshTokenRepository {
	return authrepo.NewRefreshTokenRepository(db)
}

type otpSenderAdapter struct {
	s otp.Sender
}

func NewOTPSenderAdapter(s otp.Sender) ports.OTPSender {
	return &otpSenderAdapter{s: s}
}

func (o *otpSenderAdapter) Send(ctx context.Context, to string, code string) error {
	return o.s.Send(ctx, to, code)
}
