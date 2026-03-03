package transport

import (
	"context"

	"github.com/your-moon/grape-auth-service/internal/auth/ports"
	"github.com/your-moon/grape-shared/proto/authv1"
)

const AuthServiceName = "auth"

// AuthHandler implements go-micro RPC handlers for the auth service.
type AuthHandler struct {
	svc ports.AuthService
}

func NewAuthHandler(svc ports.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

func (h *AuthHandler) LogAuth(ctx context.Context, req *authv1.LogAuthRequest, rsp *authv1.LogAuthResponse) error {
	if req != nil && req.Log != nil {
		h.svc.LogAuth(ctx, ProtoToAuthLog(req.Log))
	}
	return nil
}

func (h *AuthHandler) RequestOTP(ctx context.Context, req *authv1.IdentifierRequest, rsp *authv1.EmptyResponse) error {
	if req == nil {
		return nil
	}
	return h.svc.RequestOTP(ctx, req.Identifier)
}

func (h *AuthHandler) RequestRegistrationOTP(ctx context.Context, req *authv1.IdentifierRequest, rsp *authv1.EmptyResponse) error {
	if req == nil {
		return nil
	}
	return h.svc.RequestRegistrationOTP(ctx, req.Identifier)
}

func (h *AuthHandler) RequestForgotPasswordOTP(ctx context.Context, req *authv1.IdentifierRequest, rsp *authv1.EmptyResponse) error {
	if req == nil {
		return nil
	}
	return h.svc.RequestForgotPasswordOTP(ctx, req.Identifier)
}

func (h *AuthHandler) SignUp(ctx context.Context, req *authv1.SignUpRequest, rsp *authv1.TokenResponse) error {
	if req == nil {
		return nil
	}
	user, access, refresh, err := h.svc.SignUp(ctx, req.Identifier, req.Code, req.Password)
	if err != nil {
		code, msg := errToCode(err)
		rsp.Error = &authv1.ErrorResponse{Code: code, Message: msg}
		return nil
	}
	rsp.User = UserToProto(user)
	rsp.AccessToken = access
	rsp.RefreshToken = refresh
	return nil
}

func (h *AuthHandler) VerifyRegistration(ctx context.Context, req *authv1.VerifyRegistrationRequest, rsp *authv1.TokenResponse) error {
	if req == nil {
		return nil
	}
	user, access, refresh, err := h.svc.VerifyRegistration(ctx, req.Identifier, req.Code)
	if err != nil {
		code, msg := errToCode(err)
		rsp.Error = &authv1.ErrorResponse{Code: code, Message: msg}
		return nil
	}
	rsp.User = UserToProto(user)
	rsp.AccessToken = access
	rsp.RefreshToken = refresh
	return nil
}

func (h *AuthHandler) SetPassword(ctx context.Context, req *authv1.SetPasswordRequest, rsp *authv1.EmptyResponse) error {
	if req == nil {
		return nil
	}
	return h.svc.SetPassword(ctx, req.UserId, req.Password)
}

func (h *AuthHandler) LoginPassword(ctx context.Context, req *authv1.LoginPasswordRequest, rsp *authv1.TokenResponse) error {
	if req == nil {
		return nil
	}
	user, access, refresh, err := h.svc.LoginPassword(ctx, req.Identifier, req.Password)
	if err != nil {
		code, msg := errToCode(err)
		rsp.Error = &authv1.ErrorResponse{Code: code, Message: msg}
		return nil
	}
	rsp.User = UserToProto(user)
	rsp.AccessToken = access
	rsp.RefreshToken = refresh
	return nil
}

func (h *AuthHandler) LoginOTP(ctx context.Context, req *authv1.LoginOTPRequest, rsp *authv1.TokenResponse) error {
	if req == nil {
		return nil
	}
	user, access, refresh, err := h.svc.LoginOTP(ctx, req.Phone, req.Code)
	if err != nil {
		code, msg := errToCode(err)
		rsp.Error = &authv1.ErrorResponse{Code: code, Message: msg}
		return nil
	}
	rsp.User = UserToProto(user)
	rsp.AccessToken = access
	rsp.RefreshToken = refresh
	return nil
}

func (h *AuthHandler) ResetPassword(ctx context.Context, req *authv1.ResetPasswordRequest, rsp *authv1.EmptyResponse) error {
	if req == nil {
		return nil
	}
	return h.svc.ResetPassword(ctx, req.Identifier, req.Code, req.NewPassword)
}

func (h *AuthHandler) VerifyOTP(ctx context.Context, req *authv1.VerifyOTPRequest, rsp *authv1.EmptyResponse) error {
	if req == nil {
		return nil
	}
	return h.svc.VerifyOTP(ctx, req.Identifier, req.Code, req.Consume)
}

func (h *AuthHandler) RequestPhoneChange(ctx context.Context, req *authv1.RequestPhoneChangeRequest, rsp *authv1.EmptyResponse) error {
	if req == nil {
		return nil
	}
	return h.svc.RequestPhoneChange(ctx, req.UserId, req.NewPhone)
}

func (h *AuthHandler) ConfirmPhoneChange(ctx context.Context, req *authv1.ConfirmPhoneChangeRequest, rsp *authv1.EmptyResponse) error {
	if req == nil {
		return nil
	}
	return h.svc.ConfirmPhoneChange(ctx, req.UserId, req.NewPhone, req.Code)
}

func (h *AuthHandler) RequestEmailChange(ctx context.Context, req *authv1.RequestEmailChangeRequest, rsp *authv1.EmptyResponse) error {
	if req == nil {
		return nil
	}
	return h.svc.RequestEmailChange(ctx, req.UserId, req.NewEmail)
}

func (h *AuthHandler) ConfirmEmailChange(ctx context.Context, req *authv1.ConfirmEmailChangeRequest, rsp *authv1.EmptyResponse) error {
	if req == nil {
		return nil
	}
	return h.svc.ConfirmEmailChange(ctx, req.UserId, req.NewEmail, req.Code)
}

func (h *AuthHandler) DeleteAccount(ctx context.Context, req *authv1.DeleteAccountRequest, rsp *authv1.EmptyResponse) error {
	if req == nil {
		return nil
	}
	return h.svc.DeleteAccount(ctx, req.UserId)
}

func (h *AuthHandler) RefreshToken(ctx context.Context, req *authv1.RefreshTokenRequest, rsp *authv1.RefreshTokenResponse) error {
	if req == nil {
		return nil
	}
	access, refresh, err := h.svc.RefreshToken(ctx, req.Token)
	if err != nil {
		code, msg := errToCode(err)
		rsp.Error = &authv1.ErrorResponse{Code: code, Message: msg}
		return nil
	}
	rsp.AccessToken = access
	rsp.RefreshToken = refresh
	return nil
}

func (h *AuthHandler) Logout(ctx context.Context, req *authv1.LogoutRequest, rsp *authv1.EmptyResponse) error {
	if req == nil {
		return nil
	}
	return h.svc.Logout(ctx, req.Token)
}

func (h *AuthHandler) GetUser(ctx context.Context, req *authv1.GetUserRequest, rsp *authv1.GetUserResponse) error {
	if req == nil || rsp == nil {
		return nil
	}
	user, err := h.svc.GetUserByID(ctx, req.UserId)
	if err != nil {
		code, msg := errToCode(err)
		rsp.Error = &authv1.ErrorResponse{Code: code, Message: msg}
		return nil
	}
	rsp.User = UserToProto(user)
	return nil
}
