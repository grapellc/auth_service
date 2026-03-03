package domain

import (
	"context"
	"time"

	"github.com/your-moon/grape-shared/dto"
	"github.com/your-moon/grape-shared/entities"
	"github.com/your-moon/grape-shared/common/types"
)

type MarketService interface {
	Page(ctx context.Context, pagination types.Pagination) ([]*Market, int, error)
	GetByID(ctx context.Context, id uint) (*Market, error)
	TrackView(ctx context.Context, marketID uint, userID *uint, ipAddress *string) error
	Create(ctx context.Context, req dto.MarketCreateDTO) (*Market, error)
	Update(ctx context.Context, id uint, req dto.MarketUpdateDTO) (*Market, error)
	Delete(ctx context.Context, id uint) error
	SyncMeili()
	ToggleLike(ctx context.Context, userID uint, marketID uint) error
	IsMarketLiked(ctx context.Context, userID uint, marketID uint) (bool, error)
	GetLikeCount(ctx context.Context, marketID uint) (int64, error)
	GetLikedByUser(ctx context.Context, userID uint) ([]*Market, error)
	GetViewedByUser(ctx context.Context, userID uint) ([]*Market, error)
	// Review methods
	CreateReview(ctx context.Context, req dto.MarketReviewCreateDTO) (*MarketReview, error)
	GetReviews(ctx context.Context, marketID uint, limit, offset int) ([]*MarketReview, int64, error)
	UpdateReview(ctx context.Context, reviewID uint, userID uint, req dto.MarketReviewUpdateDTO) (*MarketReview, error)
	DeleteReview(ctx context.Context, reviewID uint, userID uint) error
	ToggleReviewLike(ctx context.Context, userID uint, reviewID uint) error
	IsReviewLiked(ctx context.Context, userID uint, reviewID uint) (bool, error)
	GetReviewLikeCount(ctx context.Context, reviewID uint) (int64, error)
	// Post methods
	CreatePost(ctx context.Context, req dto.MarketPostCreateDTO) (*MarketPost, error)
	GetPosts(ctx context.Context, marketID uint, limit, offset int) ([]*MarketPost, int64, error)
}

type Market struct {
	ID             uint                      `json:"id"`
	Name           string                    `json:"name"`
	Description    string                    `json:"description"`
	ViewCount      int64                     `json:"view_count"`
	Image          string                    `json:"image,omitempty"`
	Images         []string                  `json:"images,omitempty"`
	Location       *Location                 `json:"location,omitempty"`
	LocationDetail entities.LocationMapDetails `json:"location_details,omitempty"`
	MarketType     *MarketType                `json:"market_type,omitempty"`
	OperatingHours *entities.OperatingHours   `json:"operating_hours,omitempty"`
	CreatedAt      time.Time                 `json:"created_at"`
	UpdatedAt      time.Time                 `json:"updated_at"`
	CreatedByID    *uint                     `json:"created_by_id"`
	CreatedBy      *User                     `json:"created_by,omitempty"`
	UpdatedByID    *uint                     `json:"updated_by_id"`
	IsLiked        bool                      `json:"is_liked"`
	LikeCount      int64                     `json:"like_count"`
	Reviews        []*MarketReview           `json:"reviews,omitempty"`
	Posts          []*MarketPost             `json:"posts,omitempty"`
	Pricing        []*MarketPricing          `json:"pricing,omitempty"`
	AverageRating  float64                   `json:"average_rating"`
	ReviewCount    int64                     `json:"total_reviews"`
}

type MarketReview struct {
	ID         uint      `json:"id"`
	MarketID   uint      `json:"market_id"`
	UserID     uint      `json:"user_id"`
	Rating     int       `json:"rating"`
	Content    string    `json:"content"`
	LikesCount int64     `json:"likes_count"`
	IsLiked    bool      `json:"is_liked"`
	Images     []string  `json:"images,omitempty"`
	User       *User     `json:"user,omitempty"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type MarketPost struct {
	ID           uint      `json:"id"`
	MarketID     uint      `json:"market_id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	LikeCount    int64     `json:"like_count"`
	CommentCount int64     `json:"comment_count"`
	Images       []string  `json:"images,omitempty"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type MarketPricing struct {
	ID          uint    `json:"id"`
	MarketID    uint    `json:"market_id"`
	Name        string  `json:"name"`
	Description string  `json:"description,omitempty"`
	Tag         string  `json:"tag,omitempty"`
	Price       float64 `json:"price"`
	Image       string  `json:"image,omitempty"` // URL
}
