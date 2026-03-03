package entities

import "gorm.io/datatypes"

type AuthLog struct {
	BaseModel
	UserID        *uint          `gorm:"column:user_id;index" json:"user_id"`
	Identifier    string         `gorm:"column:identifier;index" json:"identifier"`
	Action        string         `gorm:"column:action;index" json:"action"` // login, signup, etc.
	Status        string         `gorm:"column:status;index" json:"status"` // success, failure
	FailureReason string         `gorm:"column:failure_reason" json:"failure_reason,omitempty"`
	IPAddress     string         `gorm:"column:ip_address" json:"ip_address"`
	UserAgent     string         `gorm:"column:user_agent" json:"user_agent"`
	MetaData      datatypes.JSON `gorm:"column:meta_data" json:"meta_data,omitempty"`
}

func (AuthLog) TableName() string {
	return "auth_logs"
}
