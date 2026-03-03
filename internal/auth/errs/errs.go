package errs

import "errors"

var (
	ErrUserNotFound       = errors.New("Хэрэглэгч олдсонгүй")
	ErrUserAlreadyExists  = errors.New("Энэ хэрэглэгч аль хэдийн бүртгэгдсэн байна")
	ErrPasswordNotSet     = errors.New("PASSWORD_NOT_SET")
	ErrInvalidCredentials = errors.New("Нэвтрэх мэдээлэл буруу байна")
	ErrInvalidIdentifier  = errors.New("Утасны дугаар эсвэл и-мэйл шаардлагатай")
	ErrInvalidPhone       = errors.New("Хүчингүй утасны дугаар байна")
	ErrInvalidEmail       = errors.New("Хүчингүй и-мэйл хаяг байна")
	ErrPhoneAlreadyExists = errors.New("Энэ утасны дугаар аль хэдийн бүртгэгдсэн байна")
	ErrEmailAlreadyExists = errors.New("Энэ и-мэйл аль хэдийн бүртгэгдсэн байна")

	ErrInvalidOTP          = errors.New("Баталгаажуулах код буруу байна")
	ErrOTPExpired          = errors.New("Баталгаажуулах код хүчингүй эсвэл хугацаа нь дууссан байна")
	ErrTooManyAttempts     = errors.New("Хэтэрхий олон удаа оролдсон байна. Түр хүлээнэ үү.")
	ErrOTPGenerationFailed = errors.New("Failed to generate OTP")

	ErrInvalidToken        = errors.New("Хүчингүй token байна")
	ErrInvalidRefreshToken = errors.New("Хүчингүй refresh token байна")

	ErrSystemError = errors.New("Системийн алдаа гарлаа")
)
