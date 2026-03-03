package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// RefreshTokenUUID is the RefreshToken entity with UUID primary key and user_id FK as UUID.
type RefreshTokenUUID struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	UserID    uuid.UUID      `gorm:"column:user_id;type:uuid;not null;index" json:"user_id"`
	Token     string         `gorm:"uniqueIndex;not null;size:255" json:"token"`
	ExpiresAt time.Time      `gorm:"not null" json:"expires_at"`
	Revoked   bool           `gorm:"default:false" json:"revoked"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (RefreshTokenUUID) TableName() string {
	return "refresh_tokens"
}
