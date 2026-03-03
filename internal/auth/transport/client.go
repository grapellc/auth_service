package transport

import (
	"context"

	"go-micro.dev/v5/client"

	"github.com/your-moon/grape-auth-service/internal/auth/errs"
	"github.com/your-moon/grape-auth-service/internal/auth/ports"
	"github.com/your-moon/grape-shared/entities"
	"github.com/your-moon/grape-shared/proto/authv1"
)

const endpointPrefix = "AuthHandler"

// AuthClient implements ports.AuthService by calling the auth go-micro service via RPC.
type AuthClient struct {
	c       client.Client
	svc     string
	callOpts []client.CallOption
}

// NewAuthClient returns an AuthService that calls the named auth service via the given go-micro client.
// If address is non-empty (e.g. "grape-auth-service:8060"), all calls use that address (for Docker when mdns is unavailable).
func NewAuthClient(c client.Client, serviceName string, address string) *AuthClient {
	if serviceName == "" {
		serviceName = AuthServiceName
	}
	var callOpts []client.CallOption
	if address != "" {
		callOpts = append(callOpts, client.WithAddress(address))
	}
	return &AuthClient{c: c, svc: serviceName, callOpts: callOpts}
}

func (ac *AuthClient) call(ctx context.Context, endpoint string, req, rsp interface{}) error {
	r := ac.c.NewRequest(ac.svc, endpoint, req)
	return ac.c.Call(ctx, r, rsp, ac.callOpts...)
}

func (ac *AuthClient) LogAuth(ctx context.Context, log *entities.AuthLogUUID) {
	req := &authv1.LogAuthRequest{Log: AuthLogToProto(log)}
	var rsp authv1.LogAuthResponse
	_ = ac.call(ctx, endpointPrefix+".LogAuth", req, &rsp)
}

func (ac *AuthClient) RequestOTP(ctx context.Context, identifier string) error {
	req := &authv1.IdentifierRequest{Identifier: identifier}
	var rsp authv1.EmptyResponse
	return ac.call(ctx, endpointPrefix+".RequestOTP", req, &rsp)
}

func (ac *AuthClient) RequestRegistrationOTP(ctx context.Context, identifier string) error {
	req := &authv1.IdentifierRequest{Identifier: identifier}
	var rsp authv1.EmptyResponse
	return ac.call(ctx, endpointPrefix+".RequestRegistrationOTP", req, &rsp)
}

func (ac *AuthClient) RequestForgotPasswordOTP(ctx context.Context, identifier string) error {
	req := &authv1.IdentifierRequest{Identifier: identifier}
	var rsp authv1.EmptyResponse
	return ac.call(ctx, endpointPrefix+".RequestForgotPasswordOTP", req, &rsp)
}

func (ac *AuthClient) SignUp(ctx context.Context, identifier, code, password string) (*entities.UserUUID, string, string, error) {
	req := &authv1.SignUpRequest{Identifier: identifier, Code: code, Password: password}
	var rsp authv1.TokenResponse
	if err := ac.call(ctx, endpointPrefix+".SignUp", req, &rsp); err != nil {
		return nil, "", "", err
	}
	if rsp.Error != nil {
		return nil, "", "", codeToErr(rsp.Error.Code)
	}
	return ProtoToUser(rsp.User), rsp.AccessToken, rsp.RefreshToken, nil
}

func (ac *AuthClient) VerifyRegistration(ctx context.Context, identifier, code string) (*entities.UserUUID, string, string, error) {
	req := &authv1.VerifyRegistrationRequest{Identifier: identifier, Code: code}
	var rsp authv1.TokenResponse
	if err := ac.call(ctx, endpointPrefix+".VerifyRegistration", req, &rsp); err != nil {
		return nil, "", "", err
	}
	if rsp.Error != nil {
		return nil, "", "", codeToErr(rsp.Error.Code)
	}
	return ProtoToUser(rsp.User), rsp.AccessToken, rsp.RefreshToken, nil
}

func (ac *AuthClient) SetPassword(ctx context.Context, userID string, password string) error {
	req := &authv1.SetPasswordRequest{UserId: userID, Password: password}
	var rsp authv1.EmptyResponse
	return ac.call(ctx, endpointPrefix+".SetPassword", req, &rsp)
}

func (ac *AuthClient) LoginPassword(ctx context.Context, identifier, password string) (*entities.UserUUID, string, string, error) {
	req := &authv1.LoginPasswordRequest{Identifier: identifier, Password: password}
	var rsp authv1.TokenResponse
	if err := ac.call(ctx, endpointPrefix+".LoginPassword", req, &rsp); err != nil {
		return nil, "", "", err
	}
	if rsp.Error != nil {
		return nil, "", "", codeToErr(rsp.Error.Code)
	}
	return ProtoToUser(rsp.User), rsp.AccessToken, rsp.RefreshToken, nil
}

func (ac *AuthClient) LoginOTP(ctx context.Context, phone, code string) (*entities.UserUUID, string, string, error) {
	req := &authv1.LoginOTPRequest{Phone: phone, Code: code}
	var rsp authv1.TokenResponse
	if err := ac.call(ctx, endpointPrefix+".LoginOTP", req, &rsp); err != nil {
		return nil, "", "", err
	}
	if rsp.Error != nil {
		return nil, "", "", codeToErr(rsp.Error.Code)
	}
	return ProtoToUser(rsp.User), rsp.AccessToken, rsp.RefreshToken, nil
}

func (ac *AuthClient) ResetPassword(ctx context.Context, identifier, code, newPassword string) error {
	req := &authv1.ResetPasswordRequest{Identifier: identifier, Code: code, NewPassword: newPassword}
	var rsp authv1.EmptyResponse
	return ac.call(ctx, endpointPrefix+".ResetPassword", req, &rsp)
}

func (ac *AuthClient) VerifyOTP(ctx context.Context, identifier, code string, consume bool) error {
	req := &authv1.VerifyOTPRequest{Identifier: identifier, Code: code, Consume: consume}
	var rsp authv1.EmptyResponse
	return ac.call(ctx, endpointPrefix+".VerifyOTP", req, &rsp)
}

func (ac *AuthClient) RequestPhoneChange(ctx context.Context, userID string, newPhone string) error {
	req := &authv1.RequestPhoneChangeRequest{UserId: userID, NewPhone: newPhone}
	var rsp authv1.EmptyResponse
	return ac.call(ctx, endpointPrefix+".RequestPhoneChange", req, &rsp)
}

func (ac *AuthClient) ConfirmPhoneChange(ctx context.Context, userID string, newPhone, code string) error {
	req := &authv1.ConfirmPhoneChangeRequest{UserId: userID, NewPhone: newPhone, Code: code}
	var rsp authv1.EmptyResponse
	return ac.call(ctx, endpointPrefix+".ConfirmPhoneChange", req, &rsp)
}

func (ac *AuthClient) RequestEmailChange(ctx context.Context, userID string, newEmail string) error {
	req := &authv1.RequestEmailChangeRequest{UserId: userID, NewEmail: newEmail}
	var rsp authv1.EmptyResponse
	return ac.call(ctx, endpointPrefix+".RequestEmailChange", req, &rsp)
}

func (ac *AuthClient) ConfirmEmailChange(ctx context.Context, userID string, newEmail, code string) error {
	req := &authv1.ConfirmEmailChangeRequest{UserId: userID, NewEmail: newEmail, Code: code}
	var rsp authv1.EmptyResponse
	return ac.call(ctx, endpointPrefix+".ConfirmEmailChange", req, &rsp)
}

func (ac *AuthClient) DeleteAccount(ctx context.Context, userID string) error {
	req := &authv1.DeleteAccountRequest{UserId: userID}
	var rsp authv1.EmptyResponse
	return ac.call(ctx, endpointPrefix+".DeleteAccount", req, &rsp)
}

func (ac *AuthClient) RefreshToken(ctx context.Context, token string) (string, string, error) {
	req := &authv1.RefreshTokenRequest{Token: token}
	var rsp authv1.RefreshTokenResponse
	if err := ac.call(ctx, endpointPrefix+".RefreshToken", req, &rsp); err != nil {
		return "", "", err
	}
	if rsp.Error != nil {
		return "", "", codeToErr(rsp.Error.Code)
	}
	return rsp.AccessToken, rsp.RefreshToken, nil
}

func (ac *AuthClient) Logout(ctx context.Context, token string) error {
	req := &authv1.LogoutRequest{Token: token}
	var rsp authv1.EmptyResponse
	return ac.call(ctx, endpointPrefix+".Logout", req, &rsp)
}

func (ac *AuthClient) GetUserByID(ctx context.Context, userID string) (*entities.UserUUID, error) {
	req := &authv1.GetUserRequest{UserId: userID}
	var rsp authv1.GetUserResponse
	if err := ac.call(ctx, endpointPrefix+".GetUser", req, &rsp); err != nil {
		return nil, err
	}
	if rsp.Error != nil {
		return nil, codeToErr(rsp.Error.Code)
	}
	if rsp.User == nil {
		return nil, errs.ErrUserNotFound
	}
	return ProtoToUser(rsp.User), nil
}

var _ ports.AuthService = (*AuthClient)(nil)
