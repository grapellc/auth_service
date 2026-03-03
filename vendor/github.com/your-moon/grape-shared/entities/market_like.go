package entities

type MarketLike struct {
	BaseModel
	UserID   uint    `gorm:"column:user_id;not null;uniqueIndex:idx_market_like_user_market" json:"user_id"`
	MarketID uint    `gorm:"column:market_id;not null;uniqueIndex:idx_market_like_user_market" json:"market_id"`
	User     *User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Market   *Market `gorm:"foreignKey:MarketID" json:"market,omitempty"`
}

func (MarketLike) TableName() string {
	return "market_likes"
}

type MarketReviewLike struct {
	BaseModel
	UserID   uint          `gorm:"column:user_id;not null;uniqueIndex:idx_market_review_like_user_review" json:"user_id"`
	ReviewID uint          `gorm:"column:review_id;not null;uniqueIndex:idx_market_review_like_user_review" json:"review_id"`
	User     *User         `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Review   *MarketReview `gorm:"foreignKey:ReviewID" json:"review,omitempty"`
}

func (MarketReviewLike) TableName() string {
	return "market_review_likes"
}
