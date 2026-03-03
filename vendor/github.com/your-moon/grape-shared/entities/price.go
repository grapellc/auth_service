package entities

// Price represents a product price
type Price struct {
	BaseModel
	Price     *float64 `gorm:"column:price" json:"price"`
	Currency  *string  `gorm:"column:currency" json:"currency"`
	ProductID uint     `gorm:"column:product_id;not null" json:"product_id"`

	// Relationships
	Product Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

// TableName returns the table name for Price
func (Price) TableName() string {
	return "prices"
}
