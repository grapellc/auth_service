package constants

const (
	// User validation constants
	MinEmailLength    = 5
	MaxEmailLength    = 255
	MinUsernameLength = 3
	MaxUsernameLength = 50
	MinPasswordLength = 8
	MaxPasswordLength = 128

	// Context keys
	ContextKeyUserID         = "user_id"
	OptionalContextKeyUserID = "optional_user_id"
)
