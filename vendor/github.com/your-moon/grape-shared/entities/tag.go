package entities

// Tag represents a product tag
type Tag struct {
	BaseModel
	Name string `gorm:"column:name;unique;not null" json:"name"`

	// Relationships
	Products []Product `gorm:"many2many:product_tags;" json:"products,omitempty"`
}

// TableName returns the table name for Tag
func (Tag) TableName() string {
	return "tags"
}
