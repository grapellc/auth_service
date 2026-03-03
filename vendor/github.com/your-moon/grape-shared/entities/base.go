package entities

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID          uint           `gorm:"primaryKey" json:"id"`
	CreatedAt   time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	CreatedByID *uint          `gorm:"column:created_by_id" json:"created_by_id"`
	CreatedBy   *User          `gorm:"foreignKey:CreatedByID" json:"created_by,omitempty"`
	UpdatedByID *uint          `gorm:"column:updated_by_id" json:"updated_by_id"`
	UpdatedBy   *User          `gorm:"foreignKey:UpdatedByID" json:"updated_by,omitempty"`
}

// ViewBaseModel is a base model for view tracking tables
// These tables don't have soft deletes (deleted_at) but do track updated_at
// to record when the same user/IP last viewed the item
type ViewBaseModel struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}
