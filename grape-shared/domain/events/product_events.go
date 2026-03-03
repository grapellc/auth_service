package events

type ProductViewedEvent struct {
	ProductID uint    `json:"product_id"`
	UserID    *uint   `json:"user_id,omitempty"`
	IPAddress *string `json:"ip_address,omitempty"`
}
