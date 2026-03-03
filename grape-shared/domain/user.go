package domain

import "time"

// User is used by product and graph layers (e.g. Product.CreatedBy). User CRUD lives in auth service.
type User struct {
	ID              uint      `json:"id"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
	CreatedUserID   *uint     `json:"created_user_id"`
	UpdatedUserID   *uint     `json:"updated_user_id"`
	Name            string    `json:"name,omitempty"`
	Email           string    `json:"email"`
	Username        *string   `json:"username,omitempty"`
	FirstName       *string   `json:"first_name,omitempty"`
	LastName        *string   `json:"last_name,omitempty"`
	PhoneNumber     *string   `json:"phone_number,omitempty"`
	AvatarURL       *string   `json:"avatar_url,omitempty"`
	IsPhoneVerified bool      `json:"is_phone_verified"`
	IsEmailVerified bool      `json:"is_email_verified"`
}
