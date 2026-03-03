package entities

type Product struct {
	BaseModel
	Name        string  `gorm:"column:name;not null" json:"name"`
	Description string  `gorm:"column:description" json:"description"`
	Price       float64 `gorm:"column:price" json:"price"`
	Status      string  `gorm:"column:status;default:'active'" json:"status"`
	ViewCount   int64   `gorm:"column:view_count;default:0" json:"view_count"`
	LikeCount   int64   `gorm:"column:like_count;default:0" json:"like_count"`
	IsLiked     bool    `gorm:"-" json:"is_liked"`

	CategoryID *uint     `gorm:"column:category_id" json:"category_id,omitempty"`
	Category   *Category `gorm:"foreignKey:CategoryID" json:"category,omitempty"`
	LocationID *uint     `gorm:"column:location_id" json:"location_id,omitempty"`
	Location   *Location `gorm:"foreignKey:LocationID" json:"location,omitempty"`

	// Embed LocationMapDetails directly into Product struct for flat table structure
	LocationMapDetails

	Images []*ProductImage `gorm:"foreignKey:ProductID" json:"images,omitempty"`
}

type ProductImage struct {
	ProductID uint     `gorm:"column:product_id;not null" json:"product_id"`
	FileID    uint     `gorm:"column:file_id;not null" json:"file_id"`
	Product   *Product `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	File      *File    `gorm:"foreignKey:FileID" json:"file,omitempty"`
}

// TableName returns the table name for Product
func (Product) TableName() string {
	return "products"
}

// TableName returns the table name for ProductImage
func (ProductImage) TableName() string {
	return "product_images"
}
