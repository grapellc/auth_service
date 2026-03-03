package entities

// MarketType represents a market type
type MarketType struct {
	BaseModel
	Name        string   `gorm:"column:name;not null;unique" json:"name"`
	Description string   `gorm:"column:description" json:"description"`
	Markets     []Market `gorm:"foreignKey:MarketTypeID" json:"markets,omitempty"`
}

// TableName returns the table name for MarketType
func (MarketType) TableName() string {
	return "market_types"
}
