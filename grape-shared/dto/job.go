package dto

import (
	"time"

	"github.com/your-moon/grape-shared/entities"
	"github.com/your-moon/grape-shared/common/constants"
)

type JobCreateDTO struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	JobType     constants.JobType         `json:"job_type"`
	HourlyRate  float64                   `json:"hourly_rate"`
	StartDate   *time.Time                `json:"start_date,omitempty"`
	EndDate     *time.Time                `json:"end_date,omitempty"`
	Status      constants.JobStatus       `json:"status"`
	LocationID  uint                      `json:"location_id"`
	Location    entities.LocationMapDetails `json:"location_details"`
	CompanyInfo *JobCompanyInfoDTO        `json:"company_info,omitempty"`
	ImageIDs    []uint                    `json:"image_ids"`
	UserID      uint                      `json:"-"`
}

type JobUpdateDTO struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	JobType     constants.JobType         `json:"job_type"`
	HourlyRate  float64                   `json:"hourly_rate"`
	StartDate   *time.Time                `json:"start_date,omitempty"`
	EndDate     *time.Time                `json:"end_date,omitempty"`
	Status      constants.JobStatus       `json:"status"`
	LocationID  uint                      `json:"location_id"`
	Location    entities.LocationMapDetails `json:"location_details"`
	CompanyInfo *JobCompanyInfoDTO        `json:"company_info,omitempty"`
	ImageIDs    []uint                    `json:"image_ids"`
	UserID      uint                      `json:"-"`
}

type JobCompanyInfoDTO struct {
	CompanyName string  `json:"company_name"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Salary      float64 `json:"salary"`
}

func (j *JobCreateDTO) ToModel() *entities.Job {
	uid := j.UserID
	job := &entities.Job{
		Name:               j.Name,
		Description:        j.Description,
		JobType:            j.JobType,
		HourlyRate:         j.HourlyRate,
		StartDate:          j.StartDate,
		EndDate:            j.EndDate,
		Status:             j.Status,
		LocationID:         j.LocationID,
		LocationMapDetails: j.Location,
	}
	job.CreatedByID = &uid
	job.UpdatedByID = &uid

	if j.CompanyInfo != nil {
		job.CompanyInfo = &entities.JobCompanyInfo{
			CompanyName: j.CompanyInfo.CompanyName,
			Email:       j.CompanyInfo.Email,
			Phone:       j.CompanyInfo.Phone,
			Salary:      j.CompanyInfo.Salary,
		}
	}

	return job
}

func (j *JobUpdateDTO) ToModel() *entities.Job {
	uid := j.UserID
	job := &entities.Job{
		Name:               j.Name,
		Description:        j.Description,
		JobType:            j.JobType,
		HourlyRate:         j.HourlyRate,
		StartDate:          j.StartDate,
		EndDate:            j.EndDate,
		Status:             j.Status,
		LocationID:         j.LocationID,
		LocationMapDetails: j.Location,
	}
	job.UpdatedByID = &uid

	if j.CompanyInfo != nil {
		job.CompanyInfo = &entities.JobCompanyInfo{
			CompanyName: j.CompanyInfo.CompanyName,
			Email:       j.CompanyInfo.Email,
			Phone:       j.CompanyInfo.Phone,
			Salary:      j.CompanyInfo.Salary,
		}
	}

	return job
}
