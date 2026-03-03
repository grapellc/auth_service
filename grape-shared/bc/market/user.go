package market

import "time"

type User struct {
	ID              uint      `json:"id"`
	Name            string    `json:"name,omitempty"`
	Email           string    `json:"email"`
	Username        *string   `json:"username,omitempty"`
	AvatarURL       *string   `json:"avatar_url,omitempty"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	IsPhoneVerified bool      `json:"is_phone_verified"`
	IsEmailVerified bool      `json:"is_email_verified"`
}
