package entities

// File represents a file in the system
type File struct {
	BaseModel
	Name       string  `gorm:"column:name;not null" json:"name"`
	BucketName string  `gorm:"column:bucket_name;not null" json:"bucket_name"`
	GroupName  *string `gorm:"column:group_name" json:"group_name"`
	Path       *string `gorm:"column:path" json:"path"`
	Type       *string `gorm:"column:type" json:"type"`
	Size       *int    `gorm:"column:size" json:"size"`
	URL        *string `gorm:"column:url" json:"url"`

	// Relationships
	Products []Product `gorm:"many2many:product_images;" json:"products,omitempty"`
}

// TableName returns the table name for File
func (File) TableName() string {
	return "files"
}
