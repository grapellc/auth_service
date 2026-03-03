package entities

import (
	"github.com/google/uuid"
	"gorm.io/datatypes"
)

// AuthLogUUID is the AuthLog entity with UUID primary key and user_id FK as UUID.
type AuthLogUUID struct {
	UUIDBaseModel
	UserID        *uuid.UUID     `gorm:"column:user_id;type:uuid;index" json:"user_id"`
	Identifier    string         `gorm:"column:identifier;index" json:"identifier"`
	Action        string         `gorm:"column:action;index" json:"action"`
	Status        string         `gorm:"column:status;index" json:"status"`
	FailureReason string         `gorm:"column:failure_reason" json:"failure_reason,omitempty"`
	IPAddress     string         `gorm:"column:ip_address" json:"ip_address"`
	UserAgent     string         `gorm:"column:user_agent" json:"user_agent"`
	MetaData      datatypes.JSON `gorm:"column:meta_data" json:"meta_data,omitempty"`
}

func (AuthLogUUID) TableName() string {
	return "auth_logs"
}
