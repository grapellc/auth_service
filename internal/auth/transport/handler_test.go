package transport

import (
	"context"
	"testing"

	"github.com/your-moon/grape-auth-service/internal/auth/ports"
	"github.com/your-moon/grape-shared/entities"
	"github.com/your-moon/grape-shared/proto/authv1"
	"go-micro.dev/v5/test"
)

// mockAuthService implements ports.AuthService for harness testing.
type mockAuthService struct{}

func (m *mockAuthService) LogAuth(ctx context.Context, log *entities.AuthLogUUID) {}

func (m *mockAuthService) RequestOTP(ctx context.Context, identifier string) error { return nil }
func (m *mockAuthService) RequestRegistrationOTP(ctx context.Context, identifier string) error {
	return nil
}
func (m *mockAuthService) RequestForgotPasswordOTP(ctx context.Context, identifier string) error {
	return nil
}

func (m *mockAuthService) SignUp(ctx context.Context, identifier, code, password string) (*entities.UserUUID, string, string, error) {
	return nil, "", "", nil
}
func (m *mockAuthService) VerifyRegistration(ctx context.Context, identifier, code string) (*entities.UserUUID, string, string, error) {
	return nil, "", "", nil
}
func (m *mockAuthService) SetPassword(ctx context.Context, userID string, password string) error {
	return nil
}

func (m *mockAuthService) LoginPassword(ctx context.Context, identifier, password string) (*entities.UserUUID, string, string, error) {
	return nil, "", "", nil
}
func (m *mockAuthService) LoginOTP(ctx context.Context, phone, code string) (*entities.UserUUID, string, string, error) {
	return nil, "", "", nil
}

func (m *mockAuthService) ResetPassword(ctx context.Context, identifier, code, newPassword string) error {
	return nil
}
func (m *mockAuthService) VerifyOTP(ctx context.Context, identifier, code string, consume bool) error {
	return nil
}

func (m *mockAuthService) RequestPhoneChange(ctx context.Context, userID string, newPhone string) error {
	return nil
}
func (m *mockAuthService) ConfirmPhoneChange(ctx context.Context, userID string, newPhone, code string) error {
	return nil
}
func (m *mockAuthService) RequestEmailChange(ctx context.Context, userID string, newEmail string) error {
	return nil
}
func (m *mockAuthService) ConfirmEmailChange(ctx context.Context, userID string, newEmail, code string) error {
	return nil
}

func (m *mockAuthService) DeleteAccount(ctx context.Context, userID string) error { return nil }
func (m *mockAuthService) RefreshToken(ctx context.Context, token string) (string, string, error) {
	return "", "", nil
}
func (m *mockAuthService) Logout(ctx context.Context, token string) error { return nil }

func (m *mockAuthService) GetUserByID(ctx context.Context, userID string) (*entities.UserUUID, error) {
	return nil, nil
}

var _ ports.AuthService = (*mockAuthService)(nil)

// TestAuthHandler_viaHarness verifies the auth RPC handler using go-micro test harness (see guides/testing.md).
func TestAuthHandler_viaHarness(t *testing.T) {
	h := test.NewHarness(t)
	defer h.Stop()

	handler := NewAuthHandler(&mockAuthService{})
	h.Name("auth").Register(handler)
	h.Start()

	t.Run("LogAuth", func(t *testing.T) {
		req := &authv1.LogAuthRequest{
			Log: &authv1.AuthLogMsg{
				Identifier: "test@example.com",
				Action:     "login",
				Status:     "success",
				IpAddress:  "127.0.0.1",
				UserAgent:  "test",
			},
		}
		var rsp authv1.LogAuthResponse
		err := h.Call("AuthHandler.LogAuth", req, &rsp)
		if err != nil {
			t.Fatalf("AuthHandler.LogAuth: %v", err)
		}
	})

	t.Run("RequestOTP", func(t *testing.T) {
		req := &authv1.IdentifierRequest{Identifier: "12345678"}
		var rsp authv1.EmptyResponse
		err := h.Call("AuthHandler.RequestOTP", req, &rsp)
		if err != nil {
			t.Fatalf("AuthHandler.RequestOTP: %v", err)
		}
	})
}
