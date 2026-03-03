package entities

type ProductView struct {
	ViewBaseModel
	ProductID uint     `gorm:"column:product_id;not null;index" json:"product_id"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	UserID    *uint    `gorm:"column:user_id;index" json:"user_id,omitempty"`
	User      *User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	IPAddress *string  `gorm:"column:ip_address" json:"ip_address,omitempty"`
}

// TableName returns the table name for ProductView
func (ProductView) TableName() string {
	return "product_views"
}
