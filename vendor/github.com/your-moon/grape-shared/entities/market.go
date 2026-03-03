package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
)

type OperatingHours struct {
	Monday    *DayHours `json:"monday,omitempty"`
	Tuesday   *DayHours `json:"tuesday,omitempty"`
	Wednesday *DayHours `json:"wednesday,omitempty"`
	Thursday  *DayHours `json:"thursday,omitempty"`
	Friday    *DayHours `json:"friday,omitempty"`
	Saturday  *DayHours `json:"saturday,omitempty"`
	Sunday    *DayHours `json:"sunday,omitempty"`
}

type DayHours struct {
	Open  string `json:"open"`  // Format: "09:00"
	Close string `json:"close"` // Format: "18:30"
}

// Value implements driver.Valuer interface
func (oh OperatingHours) Value() (driver.Value, error) {
	return json.Marshal(oh)
}

// Scan implements sql.Scanner interface
func (oh *OperatingHours) Scan(value interface{}) error {
	if value == nil {
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, oh)
}

type Market struct {
	BaseModel
	Creator        *User                 `gorm:"foreignKey:CreatedByID" json:"creator,omitempty"`
	Name           string                `gorm:"column:name;not null" json:"name"`
	Description    string                `gorm:"column:description" json:"description"`
	ViewCount      int64                 `gorm:"column:view_count;default:0" json:"view_count"`
	LikeCount      int64                 `gorm:"column:like_count;default:0" json:"like_count"`
	LocationID     *uint                 `gorm:"column:location_id" json:"location_id"`
	Location       *Location             `gorm:"foreignKey:LocationID" json:"location,omitempty"`
	ImageID        *uint                 `gorm:"column:image_id" json:"image_id"`
	Image          *File                 `gorm:"foreignKey:ImageID" json:"image,omitempty"`
	MarketTypeID   *uint                 `gorm:"column:market_type_id" json:"market_type_id"`
	MarketType     *MarketType           `gorm:"foreignKey:MarketTypeID" json:"market_type,omitempty"`
	Categories     []MarketCategoryModel `gorm:"many2many:market_categories;joinForeignKey:market_id;joinReferences:category_id" json:"categories,omitempty"`
	OperatingHours *OperatingHours       `gorm:"column:operating_hours;type:jsonb" json:"operating_hours,omitempty"`

	// Relationships
	Images  []*MarketImage  `gorm:"foreignKey:MarketID" json:"images,omitempty"`
	Reviews []*MarketReview `gorm:"foreignKey:MarketID" json:"reviews,omitempty"`
	Posts   []*MarketPost   `gorm:"foreignKey:MarketID" json:"posts,omitempty"`
	Pricing []MarketPricing `gorm:"foreignKey:MarketID" json:"pricing,omitempty"`

	// Embed LocationMapDetails directly into Market struct for flat table structure
	LocationMapDetails

	// Rating fields (denormalized for performance)
	AverageRating float64 `gorm:"column:average_rating;default:0" json:"average_rating"`
	ReviewCount   int64   `gorm:"column:review_count;default:0" json:"review_count"`
}

type MarketImage struct {
	MarketID uint    `gorm:"column:market_id;not null" json:"market_id"`
	FileID   uint    `gorm:"column:file_id;not null" json:"file_id"`
	Market   *Market `gorm:"foreignKey:MarketID" json:"market,omitempty"`
	File     *File   `gorm:"foreignKey:FileID" json:"file,omitempty"`
}

// TableName returns the table name for MarketImage
func (MarketImage) TableName() string {
	return "market_images"
}

type MarketReview struct {
	BaseModel
	MarketID   uint                 `gorm:"column:market_id;not null" json:"market_id"`
	UserID     uint                 `gorm:"column:user_id;not null" json:"user_id"`
	Rating     int                  `gorm:"column:rating;not null;check:rating >= 1 AND rating <= 5" json:"rating"`
	Content    string               `gorm:"column:content" json:"content"`
	LikesCount int64                `gorm:"column:likes_count;default:0" json:"likes_count"`
	Market     *Market              `gorm:"foreignKey:MarketID" json:"market,omitempty"`
	User       *User                `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Images     []*MarketReviewImage `gorm:"foreignKey:ReviewID" json:"images,omitempty"`
}

// TableName returns the table name for MarketReview
func (MarketReview) TableName() string {
	return "market_reviews"
}

type MarketReviewImage struct {
	ReviewID uint          `gorm:"column:review_id;not null" json:"review_id"`
	FileID   uint          `gorm:"column:file_id;not null" json:"file_id"`
	Review   *MarketReview `gorm:"foreignKey:ReviewID" json:"review,omitempty"`
	File     *File         `gorm:"foreignKey:FileID" json:"file,omitempty"`
}

// TableName returns the table name for MarketReviewImage
func (MarketReviewImage) TableName() string {
	return "market_review_images"
}

type MarketPost struct {
	BaseModel
	MarketID     uint               `gorm:"column:market_id;not null" json:"market_id"`
	Title        string             `gorm:"column:title;not null" json:"title"`
	Content      string             `gorm:"column:content" json:"content"`
	LikeCount    int64              `gorm:"column:like_count;default:0" json:"like_count"`
	CommentCount int64              `gorm:"column:comment_count;default:0" json:"comment_count"`
	Market       *Market            `gorm:"foreignKey:MarketID" json:"market,omitempty"`
	Images       []*MarketPostImage `gorm:"foreignKey:PostID" json:"images,omitempty"`
}

// TableName returns the table name for MarketPost
func (MarketPost) TableName() string {
	return "market_posts"
}

type MarketPostImage struct {
	PostID uint        `gorm:"column:post_id;not null" json:"post_id"`
	FileID uint        `gorm:"column:file_id;not null" json:"file_id"`
	Post   *MarketPost `gorm:"foreignKey:PostID" json:"post,omitempty"`
	File   *File       `gorm:"foreignKey:FileID" json:"file,omitempty"`
}

// TableName returns the table name for MarketPostImage
func (MarketPostImage) TableName() string {
	return "market_post_images"
}

type MarketCategoryModel struct {
	BaseModel
	Name string `gorm:"column:name;not null" json:"name"`
}

// TableName returns the table name for MarketCategoryModel
func (MarketCategoryModel) TableName() string {
	return "market_category_models"
}

type MarketCategory struct {
	BaseModel
	MarketID   uint                 `gorm:"column:market_id" json:"market_id"`
	Market     *Market              `gorm:"foreignKey:MarketID" json:"market,omitempty"`
	CategoryID uint                 `gorm:"column:category_id" json:"category_id"`
	Category   *MarketCategoryModel `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
}

type MarketPricing struct {
	BaseModel
	MarketID    uint    `gorm:"column:market_id;not null" json:"market_id"`
	Name        string  `gorm:"column:name;not null" json:"name"` // Service or Product Name
	Description string  `gorm:"column:description" json:"description"`
	Tag         string  `gorm:"column:tag" json:"tag"` // e.g., "Representative", "Popular"
	Price       float64 `gorm:"column:price;not null" json:"price"`
	ImageID     *uint   `gorm:"column:image_id" json:"image_id"`

	// Relationships
	Market *Market `gorm:"foreignKey:MarketID" json:"market,omitempty"`
	Image  *File   `gorm:"foreignKey:ImageID" json:"image,omitempty"`
}

// TableName returns the table name for MarketPricing
func (MarketPricing) TableName() string {
	return "market_pricings"
}
