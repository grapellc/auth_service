package dmo

import (
	"github.com/your-moon/grape-shared/domain"
	"github.com/your-moon/grape-shared/entities"
)

func ToJobDomain(job *entities.Job, isLiked bool, likeCount int64) *domain.Job {
	// Convert company info
	var companyInfo *domain.JobCompanyInfo
	if job.CompanyInfo != nil {
		companyInfo = &domain.JobCompanyInfo{
			ID:          job.CompanyInfo.ID,
			CompanyName: job.CompanyInfo.CompanyName,
			Email:       job.CompanyInfo.Email,
			Phone:       job.CompanyInfo.Phone,
			Salary:      job.CompanyInfo.Salary,
		}
	}

	// Convert location
	var location *domain.Location
	if job.Location != nil {
		location = ToLocation(job.Location)
	} else if job.Latitude != nil && job.Longitude != nil {
		// Fallback to coordinates if Location relation is missing
		name := "Сонгосон байршил"
		if job.Locality != nil && *job.Locality != "" {
			name = *job.Locality
		} else if job.Region != nil && *job.Region != "" {
			name = *job.Region
		}

		location = &domain.Location{
			Name: name,
			Lat:  *job.Latitude,
			Lng:  *job.Longitude,
		}
	}

	// Convert images to string URLs
	var imageURLs []string
	var thumbnail string
	if job.Images != nil {
		for i, img := range job.Images {
			if img.File != nil && img.File.URL != nil {
				imageURLs = append(imageURLs, *img.File.URL)
				// Set thumbnail to the first image
				if i == 0 {
					thumbnail = *img.File.URL
				}
			}
		}
	}

	// Convert user
	var user *domain.User
	if job.Creator != nil {
		user = ToUserDomain(job.Creator)
	} else if job.CreatedBy != nil {
		user = ToUserDomain(job.CreatedBy)
	}

	return &domain.Job{
		ID:             job.ID,
		Name:           job.Name,
		Description:    job.Description,
		JobType:        job.JobType,
		HourlyRate:     job.HourlyRate,
		StartDate:      job.StartDate,
		EndDate:        job.EndDate,
		Status:         job.Status,
		ViewCount:      job.ViewCount,
		Location:       location,
		LocationDetail: job.LocationMapDetails,
		CompanyInfo:    companyInfo,
		Images:         imageURLs,
		Thumbnail:      thumbnail,
		CreatedAt:      job.CreatedAt,
		UpdatedAt:      job.UpdatedAt,
		CreatedByID:    job.CreatedByID,
		CreatedBy:      user,
		UpdatedByID:    job.UpdatedByID,
		IsLiked:        isLiked,
		LikeCount:      likeCount,
	}
}
