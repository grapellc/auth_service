package entities

import "github.com/google/uuid"

// UserUUID is the User entity with UUID primary key.
// Use this when the users table is created with uuid id (see 001_auth_tables_uuid.sql).
type UserUUID struct {
	UUIDBaseModelNoAudit
	CreatedUserID *uuid.UUID `gorm:"column:created_user_id;type:uuid" json:"created_user_id"`
	UpdatedUserID *uuid.UUID `gorm:"column:updated_user_id;type:uuid" json:"updated_user_id"`
	ClerkID       *string    `gorm:"column:clerk_id" json:"clerk_id"`
	Email         string     `gorm:"column:email;not null" json:"email"`
	Username      *string    `gorm:"column:username" json:"username"`
	FirstName     *string    `gorm:"column:first_name" json:"first_name"`
	LastName      *string    `gorm:"column:last_name" json:"last_name"`
	PhoneNumber   *string    `gorm:"column:phone_number" json:"phone_number"`
	AvatarURL     *string    `gorm:"column:avatar_url" json:"avatar_url"`
	PasswordHash  *string    `gorm:"column:password_hash" json:"-"`
	Role          string     `gorm:"column:role;default:user" json:"role"`
	IsPhoneVerified bool     `gorm:"column:is_phone_verified;default:false" json:"is_phone_verified"`
	IsEmailVerified bool     `gorm:"column:is_email_verified;default:false" json:"is_email_verified"`
}

func (UserUUID) TableName() string {
	return "users"
}
