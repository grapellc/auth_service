package entities

type JobView struct {
	ViewBaseModel
	JobID     uint    `gorm:"column:job_id;not null;index" json:"job_id"`
	Job       *Job    `gorm:"foreignKey:JobID" json:"job,omitempty"`
	UserID    *uint   `gorm:"column:user_id;index" json:"user_id,omitempty"`
	User      *User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	IPAddress *string `gorm:"column:ip_address" json:"ip_address,omitempty"`
}

// TableName returns the table name for JobView
func (JobView) TableName() string {
	return "job_views"
}
