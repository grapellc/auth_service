package entities

type SystemSetting struct {
	BaseModel
	Key   string `gorm:"column:key;uniqueIndex;not null" json:"key"`
	Value string `gorm:"column:value;not null" json:"value"`
}

func (SystemSetting) TableName() string {
	return "system_settings"
}
