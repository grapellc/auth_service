package entities

type ProductLike struct {
	BaseModel
	UserID    uint     `gorm:"column:user_id;not null;uniqueIndex:idx_product_like_user_product" json:"user_id"`
	ProductID uint     `gorm:"column:product_id;not null;uniqueIndex:idx_product_like_user_product" json:"product_id"`
	User      *User    `gorm:"foreignKey:UserID" json:"user,omitempty"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

func (ProductLike) TableName() string {
	return "product_likes"
}
