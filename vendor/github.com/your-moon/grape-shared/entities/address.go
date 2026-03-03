package entities

// Address represents an address
type Address struct {
	BaseModel
	Name      string   `gorm:"column:name;not null" json:"name"`
	Address   *string  `gorm:"column:address" json:"address"`
	Latitude  *float64 `gorm:"column:latitude" json:"latitude"`
	Longitude *float64 `gorm:"column:longitude" json:"longitude"`
}

// TableName returns the table name for Address
func (Address) TableName() string {
	return "addresses"
}
