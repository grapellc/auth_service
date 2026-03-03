package entities

import (
	"time"

	"github.com/your-moon/grape-shared/common/constants"
)

type Job struct {
	BaseModel
	Creator     *User               `gorm:"foreignKey:CreatedByID" json:"creator,omitempty"`
	Name        string              `gorm:"column:name;not null" json:"name"`
	Description string              `gorm:"column:description" json:"description"`
	JobType     constants.JobType   `gorm:"column:job_type" json:"job_type"`
	HourlyRate  float64             `gorm:"column:hourly_rate" json:"hourly_rate"`
	StartDate   *time.Time          `gorm:"column:start_date" json:"start_date"`
	EndDate     *time.Time          `gorm:"column:end_date" json:"end_date"`
	Status      constants.JobStatus `gorm:"column:status" json:"status"`
	ViewCount   int64               `gorm:"column:view_count;default:0" json:"view_count"`
	LikeCount   int64               `gorm:"column:like_count;default:0" json:"like_count"`
	LocationID  uint                `gorm:"column:location_id" json:"location_id"`
	Location    *Location           `gorm:"foreignKey:LocationID" json:"location,omitempty"`
	CompanyInfo *JobCompanyInfo     `gorm:"foreignKey:JobID" json:"company_info,omitempty"`

	// Embed LocationMapDetails directly into Job struct for flat table structure
	LocationMapDetails

	Images []*JobImage `gorm:"foreignKey:JobID" json:"images,omitempty"`
}

type JobImage struct {
	JobID  uint  `gorm:"column:job_id;not null" json:"job_id"`
	FileID uint  `gorm:"column:file_id;not null" json:"file_id"`
	Job    *Job  `gorm:"foreignKey:JobID" json:"job,omitempty"`
	File   *File `gorm:"foreignKey:FileID" json:"file,omitempty"`
}

// TableName returns the table name for JobImage
func (JobImage) TableName() string {
	return "job_images"
}

type JobCompanyInfo struct {
	BaseModel
	JobID       uint    `gorm:"column:job_id;not null" json:"job_id"`
	CompanyName string  `gorm:"column:company_name;not null" json:"company_name"`
	Email       string  `gorm:"column:email;not null" json:"email"`
	Phone       string  `gorm:"column:phone" json:"phone"`
	Salary      float64 `gorm:"column:salary" json:"salary"`
}

type JobCategory struct {
	BaseModel
	Name string `gorm:"column:name;not null" json:"name"`
}
