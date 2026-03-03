package entities

type MarketView struct {
	ViewBaseModel
	MarketID  uint    `gorm:"column:market_id;not null;index" json:"market_id"`
	Market    *Market `gorm:"foreignKey:MarketID" json:"market,omitempty"`
	UserID    *uint   `gorm:"column:user_id;index" json:"user_id,omitempty"`
	User      *User   `gorm:"foreignKey:UserID" json:"user,omitempty"`
	IPAddress *string `gorm:"column:ip_address" json:"ip_address,omitempty"`
}

// TableName returns the table name for MarketView
func (MarketView) TableName() string {
	return "market_views"
}
