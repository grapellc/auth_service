package events

type MarketViewedEvent struct {
	MarketID  uint    `json:"market_id"`
	UserID    *uint   `json:"user_id,omitempty"`
	IPAddress *string `json:"ip_address,omitempty"`
}

type ReviewCreatedEvent struct {
	ReviewID uint `json:"review_id"`
	MarketID uint `json:"market_id"`
	Rating   int  `json:"rating"`
}

type ReviewUpdatedEvent struct {
	ReviewID uint `json:"review_id"`
	MarketID uint `json:"market_id"`
	Rating   int  `json:"rating"`
}

type ReviewDeletedEvent struct {
	ReviewID uint `json:"review_id"`
	MarketID uint `json:"market_id"`
}
