package entities

// Module represents a system module
type Module struct {
	BaseModel
	Name         string  `gorm:"column:name;not null" json:"name"`
	Slug         *string `gorm:"column:slug" json:"slug"`
	Description  *string `gorm:"column:description" json:"description"`
	IconLink     *string `gorm:"column:icon_link" json:"icon_link"`
	URL          *string `gorm:"column:url" json:"url"`
	Module       *string `gorm:"column:module" json:"module"`
	Status       string  `gorm:"column:status;default:active" json:"status"`
	MobileStatus string  `gorm:"column:mobile_status;default:active" json:"mobile_status"`
	Index        int     `gorm:"column:index;default:0" json:"index"`
}

// TableName returns the table name for Module
func (Module) TableName() string {
	return "modules"
}
