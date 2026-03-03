package entities

// AdvertisementWords represents advertisement words
type AdvertisementWords struct {
	BaseModel
	Word *string `gorm:"column:word" json:"word"`
}

// TableName returns the table name for AdvertisementWords
func (AdvertisementWords) TableName() string {
	return "advertisement_words"
}
