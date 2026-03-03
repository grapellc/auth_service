package entities

// Category represents a product category
type Category struct {
	BaseModel
	Name string `gorm:"column:name;not null" json:"name"`

	// Relationships
	Products []Product `gorm:"foreignKey:CategoryID" json:"products,omitempty"`
}

// TableName returns the table name for Category
func (Category) TableName() string {
	return "categories"
}
