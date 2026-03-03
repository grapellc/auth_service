package entities

import "time"

type SearchLog struct {
	BaseModel
	Query      string `gorm:"column:query;index" json:"query"`
	ModuleSlug string `gorm:"column:module_slug;index" json:"module_slug"`
}

// TableName returns the table name for SearchLog
func (SearchLog) TableName() string {
	return "search_logs"
}

type SearchStat struct {
	Query     string    `gorm:"primaryKey;column:query" json:"query"`
	Score     int64     `gorm:"column:score;index" json:"score"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (SearchStat) TableName() string {
	return "search_stats"
}
