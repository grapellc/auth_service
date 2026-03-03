package entities

import (
	"time"

	"gorm.io/gorm"
)

type UserBaseModel struct {
	ID            uint           `gorm:"primaryKey" json:"id"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	CreatedUserID *uint          `gorm:"column:created_user_id" json:"created_user_id"`
	UpdatedUserID *uint          `gorm:"column:updated_user_id" json:"updated_user_id"`
}

// User represents a user in the system
type User struct {
	UserBaseModel
	Email           string  `gorm:"column:email;not null" json:"email"`
	Username        *string `gorm:"column:username" json:"username"`
	FirstName       *string `gorm:"column:first_name" json:"first_name"`
	LastName        *string `gorm:"column:last_name" json:"last_name"`
	PhoneNumber     *string `gorm:"column:phone_number" json:"phone_number"`
	AvatarURL       *string `gorm:"column:avatar_url" json:"avatar_url"`
	PasswordHash    *string `gorm:"column:password_hash" json:"-"`
	Role            string  `gorm:"column:role;default:user" json:"role"`
	IsPhoneVerified bool    `gorm:"column:is_phone_verified;default:false" json:"is_phone_verified"`
	IsEmailVerified bool    `gorm:"column:is_email_verified;default:false" json:"is_email_verified"`

	// Relationships
	CreatedProducts []Product `gorm:"foreignKey:CreatedByID" json:"created_products,omitempty"`
	UpdatedProducts []Product `gorm:"foreignKey:UpdatedByID" json:"updated_products,omitempty"`
}

// TableName returns the table name for User
func (User) TableName() string {
	return "users"
}
