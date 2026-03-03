package market

import (
	"time"

	"github.com/your-moon/grape-shared/entities"
)

type Market struct {
	ID             uint                       `json:"id"`
	Name           string                     `json:"name"`
	Description    string                     `json:"description"`
	ViewCount      int64                      `json:"view_count"`
	Image          string                     `json:"image,omitempty"`
	Images         []string                   `json:"images,omitempty"`
	Location       *Location                  `json:"location,omitempty"`
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
	Image       string  `json:"image,omitempty"`
}
