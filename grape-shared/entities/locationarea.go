package entities

// LocationArea represents a location area
type LocationArea struct {
	BaseModel
	Name string `gorm:"column:name;not null" json:"name"`

	// Relationships
	Locations []Location `gorm:"foreignKey:LocationAreaID" json:"locations,omitempty"`
}

// TableName returns the table name for LocationArea
func (LocationArea) TableName() string {
	return "location_areas"
}
