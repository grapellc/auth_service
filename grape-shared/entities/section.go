package entities

// Section represents a system section
type Section struct {
	BaseModel
	Name        string  `gorm:"column:name;not null" json:"name"`
	Slug        *string `gorm:"column:slug" json:"slug"`
	Description *string `gorm:"column:description" json:"description"`
	IconLink    *string `gorm:"column:icon_link" json:"icon_link"`
	URL         *string `gorm:"column:url" json:"url"`
	Status      string  `gorm:"column:status;default:active" json:"status"`
	Index       int     `gorm:"column:index;default:0" json:"index"`
}

// TableName returns the table name for Section
func (Section) TableName() string {
	return "sections"
}
