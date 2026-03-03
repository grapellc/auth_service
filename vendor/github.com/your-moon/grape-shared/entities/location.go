package entities

// Location represents a location
type Location struct {
	BaseModel
	Name           string  `gorm:"column:name" json:"name"`
	Address        string  `gorm:"column:address" json:"address"`
	Latitude       float64 `gorm:"column:latitude" json:"latitude"`
	Longitude      float64 `gorm:"column:longitude" json:"longitude"`
	LocationAreaID *uint   `gorm:"column:location_area_id" json:"location_area_id"`

	// Relationships
	LocationArea *LocationArea `gorm:"foreignKey:LocationAreaID" json:"location_area,omitempty"`
	Products     []Product     `gorm:"foreignKey:LocationID" json:"products,omitempty"`
}

func (loc *Location) ToDomain() interface{} {
	return struct {
		ID      uint    `json:"id"`
		Name    string  `json:"name"`
		Address string  `json:"address"`
		Lat     float64 `json:"lat"`
		Lng     float64 `json:"lng"`
	}{
		ID:      loc.ID,
		Name:    loc.Name,
		Address: loc.Address,
		Lat:     loc.Latitude,
		Lng:     loc.Longitude,
	}
}

// TableName returns the table name for Location
func (Location) TableName() string {
	return "locations"
}
