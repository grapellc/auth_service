package entities

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// UUIDBaseModel is a base model with UUID primary key for entities that use UUIDs.
type UUIDBaseModel struct {
	ID            uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CreatedAt     time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt     time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt     gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
	CreatedByID   *uuid.UUID     `gorm:"column:created_by_id;type:uuid" json:"created_by_id"`
	UpdatedByID   *uuid.UUID     `gorm:"column:updated_by_id;type:uuid" json:"updated_by_id"`
}

// UUIDBaseModelNoAudit is a base with UUID primary key and no created_by/updated_by.
type UUIDBaseModelNoAudit struct {
	ID        uuid.UUID      `gorm:"type:uuid;primaryKey;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;index" json:"-"`
}
