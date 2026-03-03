package transport

import (
	"errors"

	"github.com/your-moon/grape-auth-service/internal/auth/errs"
)

// Error codes for RPC (must match client mapping)
const (
	CodeUserNotFound       = "USER_NOT_FOUND"
	CodeUserAlreadyExists  = "USER_ALREADY_EXISTS"
	CodePasswordNotSet     = "PASSWORD_NOT_SET"
	CodeInvalidCredentials = "INVALID_CREDENTIALS"
	CodeInvalidIdentifier  = "INVALID_IDENTIFIER"
	CodeInvalidEmail       = "INVALID_EMAIL"
	CodePhoneAlreadyExists = "PHONE_ALREADY_EXISTS"
	CodeEmailAlreadyExists = "EMAIL_ALREADY_EXISTS"
	CodeInvalidOTP         = "INVALID_OTP"
	CodeOTPExpired         = "OTP_EXPIRED"
	CodeTooManyAttempts    = "TOO_MANY_ATTEMPTS"
	CodeInvalidRefreshToken = "INVALID_REFRESH_TOKEN"
)

func errToCode(err error) (code, message string) {
	if err == nil {
		return "", ""
	}
	message = err.Error()
	switch {
	case errors.Is(err, errs.ErrUserNotFound):
		return CodeUserNotFound, message
	case errors.Is(err, errs.ErrUserAlreadyExists):
		return CodeUserAlreadyExists, message
	case errors.Is(err, errs.ErrPasswordNotSet):
		return CodePasswordNotSet, message
	case errors.Is(err, errs.ErrInvalidCredentials):
		return CodeInvalidCredentials, message
	case errors.Is(err, errs.ErrInvalidIdentifier):
		return CodeInvalidIdentifier, message
	case errors.Is(err, errs.ErrInvalidEmail):
		return CodeInvalidEmail, message
	case errors.Is(err, errs.ErrPhoneAlreadyExists):
		return CodePhoneAlreadyExists, message
	case errors.Is(err, errs.ErrEmailAlreadyExists):
		return CodeEmailAlreadyExists, message
	case errors.Is(err, errs.ErrInvalidOTP):
		return CodeInvalidOTP, message
	case errors.Is(err, errs.ErrOTPExpired):
		return CodeOTPExpired, message
	case errors.Is(err, errs.ErrTooManyAttempts):
		return CodeTooManyAttempts, message
	case errors.Is(err, errs.ErrInvalidRefreshToken):
		return CodeInvalidRefreshToken, message
	default:
		return "UNKNOWN", message
	}
}

func codeToErr(code string) error {
	switch code {
	case CodeUserNotFound:
		return errs.ErrUserNotFound
	case CodeUserAlreadyExists:
		return errs.ErrUserAlreadyExists
	case CodePasswordNotSet:
		return errs.ErrPasswordNotSet
	case CodeInvalidCredentials:
		return errs.ErrInvalidCredentials
	case CodeInvalidIdentifier:
		return errs.ErrInvalidIdentifier
	case CodeInvalidEmail:
		return errs.ErrInvalidEmail
	case CodePhoneAlreadyExists:
		return errs.ErrPhoneAlreadyExists
	case CodeEmailAlreadyExists:
		return errs.ErrEmailAlreadyExists
	case CodeInvalidOTP:
		return errs.ErrInvalidOTP
	case CodeOTPExpired:
		return errs.ErrOTPExpired
	case CodeTooManyAttempts:
		return errs.ErrTooManyAttempts
	case CodeInvalidRefreshToken:
		return errs.ErrInvalidRefreshToken
	default:
		return errs.ErrSystemError
	}
}
