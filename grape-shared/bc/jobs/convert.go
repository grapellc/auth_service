package jobs

import (
	"github.com/your-moon/grape-shared/domain"
)

// FromDomainJob converts shared domain.Job to jobs.Job.
func FromDomainJob(d *domain.Job) *Job {
	if d == nil {
		return nil
	}
	var companyInfo *JobCompanyInfo
	if d.CompanyInfo != nil {
		companyInfo = &JobCompanyInfo{
			ID:          d.CompanyInfo.ID,
			CompanyName: d.CompanyInfo.CompanyName,
			Email:       d.CompanyInfo.Email,
			Phone:       d.CompanyInfo.Phone,
			Salary:      d.CompanyInfo.Salary,
		}
	}
	return &Job{
		ID:             d.ID,
		Name:           d.Name,
		Description:    d.Description,
		JobType:        d.JobType,
		HourlyRate:     d.HourlyRate,
		StartDate:      d.StartDate,
		EndDate:        d.EndDate,
		Status:         d.Status,
		ViewCount:      d.ViewCount,
		Location:       d.Location,
		LocationDetail: d.LocationDetail,
		CompanyInfo:    companyInfo,
		Images:         d.Images,
		Thumbnail:      d.Thumbnail,
		CreatedAt:      d.CreatedAt,
		UpdatedAt:      d.UpdatedAt,
		CreatedByID:    d.CreatedByID,
		CreatedBy:      d.CreatedBy,
		UpdatedByID:    d.UpdatedByID,
		IsLiked:        d.IsLiked,
		LikeCount:      d.LikeCount,
	}
}
