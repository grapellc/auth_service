package product

import (
	"time"

	"github.com/your-moon/grape-shared/domain"
	"github.com/your-moon/grape-shared/entities"
)

// Product is the product aggregate root in the product bounded context.
type Product struct {
	ID             uint                      `json:"id"`
	Name           string                    `json:"name"`
	Description    string                    `json:"description"`
	Price          float64                   `json:"price"`
	ViewCount      int64                     `json:"view_count"`
	Status         string                    `json:"status"`
	Category       *Category                 `json:"category,omitempty"`
	Location       *domain.Location          `json:"location,omitempty"`
	LocationDetail entities.LocationMapDetails `json:"location_details,omitempty"`
	Images         []string                  `json:"images,omitempty"`
	Thumbnail      string                    `json:"thumbnail,omitempty"`
	CreatedAt      time.Time                 `json:"created_at"`
	UpdatedAt      time.Time                 `json:"updated_at"`
	CreatedByID    *uint                     `json:"created_by_id"`
	CreatedBy      *domain.User              `json:"created_by,omitempty"`
	UpdatedByID    *uint                     `json:"updated_by_id"`
	IsLiked        bool                      `json:"is_liked"`
	LikeCount      int64                     `json:"like_count"`
}

// Category is a product category in the product bounded context.
type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
