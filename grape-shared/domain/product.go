package domain

import (
	"context"
	"time"

	"github.com/your-moon/grape-shared/dto"
	"github.com/your-moon/grape-shared/entities"
	"github.com/your-moon/grape-shared/common/types"
)

type ProductService interface {
	Page(ctx context.Context, pagination types.ProductPagination) ([]*Product, int, error)
	GetByID(ctx context.Context, id uint) (*Product, error)
	GetLikeCount(ctx context.Context, productID uint) (int64, error)
	IsProductLiked(ctx context.Context, userID uint, productID uint) (bool, error)
	GetLikedByUser(ctx context.Context, userID uint) ([]*Product, error)
	TrackView(ctx context.Context, productID uint, userID *uint, ipAddress *string) error
	Create(ctx context.Context, req dto.ProductCreateDTO) (*Product, error)
	Update(ctx context.Context, id uint, req dto.ProductUpdateDTO) (*Product, error)
	Delete(ctx context.Context, id uint) error
	ToggleLike(ctx context.Context, userID uint, productID uint) (bool, error)
	GetViewedByUser(ctx context.Context, userID uint) ([]*Product, error)
	SyncMeili()
}

type Product struct {
	ID             uint                      `json:"id"`
	Name           string                    `json:"name"`
	Description    string                    `json:"description"`
	Price          float64                   `json:"price"`
	ViewCount      int64                     `json:"view_count"`
	Status         string                    `json:"status"`
	Category       *Category                 `json:"category,omitempty"`
	Location       *Location                 `json:"location,omitempty"`
	LocationDetail entities.LocationMapDetails `json:"location_details,omitempty"`
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

type Category struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
