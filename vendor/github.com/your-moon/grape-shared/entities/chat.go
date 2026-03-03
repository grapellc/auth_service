package entities

import "time"

// Conversation represents a chat conversation between two users
type Conversation struct {
	BaseModel
	User1ID uint `gorm:"not null;index" json:"user1_id"`
	User1   User `gorm:"foreignKey:User1ID" json:"user1"`
	User2ID uint `gorm:"not null;index" json:"user2_id"`
	User2   User `gorm:"foreignKey:User2ID" json:"user2"`
	// Generic Context
	ContextType string `gorm:"index" json:"context_type"` // e.g. "product", "job"
	ContextID   uint   `gorm:"index" json:"context_id"`

	// Associations for generic context (Manual Fetch)
	Job    *Job    `gorm:"-" json:"job,omitempty"`
	Market *Market `gorm:"-" json:"market,omitempty"`

	// Deprecated: Use ContextType="product" and ContextID
	ProductID      *uint      `gorm:"index" json:"product_id"`
	Product        *Product   `gorm:"foreignKey:ProductID" json:"product"`
	LastMessage    string     `json:"last_message"`
	LastMessageAt  *time.Time `json:"last_message_at"`
	DeletedByUser1 bool       `gorm:"default:false" json:"deleted_by_user1"`
	DeletedByUser2 bool       `gorm:"default:false" json:"deleted_by_user2"`
	PinnedByUser1  bool       `gorm:"default:false" json:"pinned_by_user1"`
	PinnedByUser2  bool       `gorm:"default:false" json:"pinned_by_user2"`
}

// Message represents a single message in a conversation
type Message struct {
	BaseModel
	ConversationID uint         `gorm:"not null;index" json:"conversation_id"`
	Conversation   Conversation `gorm:"foreignKey:ConversationID" json:"-"`
	SenderID       uint         `gorm:"not null;index" json:"sender_id"`
	Sender         User         `gorm:"foreignKey:SenderID" json:"sender"`
	Content        string       `gorm:"type:text;not null" json:"content"`
	IsRead         bool         `gorm:"default:false" json:"is_read"`
}
