package domain

import (
	"context"
	"time"

	"github.com/your-moon/grape-shared/dto"
	"github.com/your-moon/grape-shared/entities"
	"github.com/your-moon/grape-shared/common/constants"
	"github.com/your-moon/grape-shared/common/types"
)

type JobService interface {
	Page(ctx context.Context, pagination types.JobPagination) ([]*Job, int, error)
	GetByID(ctx context.Context, id uint) (*Job, error)
	TrackView(ctx context.Context, jobID uint, userID *uint, ipAddress *string) error
	Create(ctx context.Context, req dto.JobCreateDTO) (*Job, error)
	Update(ctx context.Context, id uint, req dto.JobUpdateDTO) (*Job, error)
	Delete(ctx context.Context, id uint) error
	SyncMeili()
	ToggleLike(ctx context.Context, userID uint, jobID uint) error
	IsJobLiked(ctx context.Context, userID uint, jobID uint) (bool, error)
	GetLikeCount(ctx context.Context, jobID uint) (int64, error)
	GetLikedByUser(ctx context.Context, userID uint) ([]*Job, error)
	GetViewedByUser(ctx context.Context, userID uint) ([]*Job, error)
}

type Job struct {
	ID             uint                      `json:"id"`
	Name           string                    `json:"name"`
	Description    string                    `json:"description"`
	JobType        constants.JobType         `json:"job_type"`
	HourlyRate     float64                   `json:"hourly_rate"`
	StartDate      *time.Time                `json:"start_date,omitempty"`
	EndDate        *time.Time                `json:"end_date,omitempty"`
	Status         constants.JobStatus       `json:"status"`
	ViewCount      int64                     `json:"view_count"`
	Location       *Location                 `json:"location,omitempty"`
	LocationDetail entities.LocationMapDetails `json:"location_details,omitempty"`
	CompanyInfo    *JobCompanyInfo           `json:"company_info,omitempty"`
	Images         []string                  `json:"images,omitempty"`
	Thumbnail      string                    `json:"thumbnail,omitempty"`
	CreatedAt      time.Time                 `json:"created_at"`
	UpdatedAt      time.Time                 `json:"updated_at"`
	CreatedByID    *uint                     `json:"created_by_id"`
	CreatedBy      *User                     `json:"created_by,omitempty"`
	UpdatedByID    *uint                     `json:"updated_by_id"`
	IsLiked        bool                      `json:"is_liked"`
	LikeCount      int64                     `json:"like_count"`
}

type JobCompanyInfo struct {
	ID          uint    `json:"id"`
	CompanyName string  `json:"company_name"`
	Email       string  `json:"email"`
	Phone       string  `json:"phone"`
	Salary      float64 `json:"salary"`
}
