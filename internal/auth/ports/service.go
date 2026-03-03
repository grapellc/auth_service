package ports

import (
	"context"

	"github.com/your-moon/grape-shared/entities"
)

// AuthService is the port for auth use cases (UUID schema).
type AuthService interface {
	LogAuth(ctx context.Context, log *entities.AuthLogUUID)

	RequestOTP(ctx context.Context, identifier string) error
	RequestRegistrationOTP(ctx context.Context, identifier string) error
	RequestForgotPasswordOTP(ctx context.Context, identifier string) error

	SignUp(ctx context.Context, identifier, code, password string) (*entities.UserUUID, string, string, error)
	VerifyRegistration(ctx context.Context, identifier, code string) (*entities.UserUUID, string, string, error)
	SetPassword(ctx context.Context, userID string, password string) error

	LoginPassword(ctx context.Context, identifier, password string) (*entities.UserUUID, string, string, error)
	LoginOTP(ctx context.Context, phone, code string) (*entities.UserUUID, string, string, error)

	ResetPassword(ctx context.Context, identifier, code, newPassword string) error
	VerifyOTP(ctx context.Context, identifier, code string, consume bool) error

	RequestPhoneChange(ctx context.Context, userID string, newPhone string) error
	ConfirmPhoneChange(ctx context.Context, userID string, newPhone, code string) error
	RequestEmailChange(ctx context.Context, userID string, newEmail string) error
	ConfirmEmailChange(ctx context.Context, userID string, newEmail, code string) error

	DeleteAccount(ctx context.Context, userID string) error
	RefreshToken(ctx context.Context, token string) (string, string, error)
	Logout(ctx context.Context, token string) error

	GetUserByID(ctx context.Context, userID string) (*entities.UserUUID, error)
}
