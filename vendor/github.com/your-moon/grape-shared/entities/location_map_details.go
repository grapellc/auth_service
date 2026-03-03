package entities

// LocationMapDetails represents specific location information usually retrieved from geocoding
type LocationMapDetails struct {
	Latitude  *float64 `gorm:"column:latitude" json:"latitude"`
	Longitude *float64 `gorm:"column:longitude" json:"longitude"`
	Postcode  *string  `gorm:"column:postcode" json:"postcode"`
	Locality  *string  `gorm:"column:locality" json:"locality"`
	Region    *string  `gorm:"column:region" json:"region"`
	Country   *string  `gorm:"column:country" json:"country"`
}
