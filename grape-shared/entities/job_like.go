package entities

type JobLike struct {
	BaseModel
	UserID uint  `gorm:"column:user_id;not null;uniqueIndex:idx_job_like_user_job" json:"user_id"`
	JobID  uint  `gorm:"column:job_id;not null;uniqueIndex:idx_job_like_user_job" json:"job_id"`
	User   *User `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Job    *Job  `gorm:"foreignKey:JobID" json:"job,omitempty"`
}

func (JobLike) TableName() string {
	return "job_likes"
}
