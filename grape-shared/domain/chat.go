package domain

import (
	"context"
	"time"

	"github.com/your-moon/grape-shared/dto"
)

type Conversation struct {
	ID            uint                 `json:"id"`
	User          *User                `json:"user"` // The other user
	LastMessage   string               `json:"last_message"`
	LastMessageAt *time.Time           `json:"last_message_at"`
	UnreadCount   int                  `json:"unread_count"`
	IsBuying      bool                 `json:"is_buying"`
	IsSelling     bool                 `json:"is_selling"`
	IsPinned      bool                 `json:"is_pinned"`
	Context       *ConversationContext `json:"context,omitempty"`
}

type ConversationContext struct {
	ID       uint   `json:"id"`
	Type     string `json:"type"` // "product", "job", "market"
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
	ImageURL string `json:"image_url"`
	Status   string `json:"status,omitempty"`
}

type Message struct {
	ID             uint      `json:"id"`
	ConversationID uint      `json:"conversation_id"`
	SenderID       uint      `json:"sender_id"`
	Sender         *User     `json:"sender"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"created_at"`
	IsRead         bool      `json:"is_read"`
}

type ChatService interface {
	CreateConversation(ctx context.Context, userID uint, req dto.CreateConversationDTO) (*Conversation, error)
	GetUserConversations(ctx context.Context, userID uint) ([]Conversation, error)
	GetConversation(ctx context.Context, conversationID uint, userID uint) (*Conversation, error)
	CreateMessage(ctx context.Context, userID uint, req dto.CreateMessageDTO) (*Message, error)
	GetMessages(ctx context.Context, conversationID uint, userID uint, limit, offset int) ([]Message, error)
	MarkAsRead(ctx context.Context, conversationID uint, userID uint) error
	DeleteConversation(ctx context.Context, conversationID uint, userID uint) error
	UpdatePinStatus(ctx context.Context, conversationID uint, userID uint, isPinned bool) error
	GetTotalUnreadCount(ctx context.Context, userID uint) (int64, error)
	GetConversationByContext(ctx context.Context, userID uint, targetUserID uint, contextType string, contextID uint) (*Conversation, error)
	RegisterClient(userID uint, conn interface{})
	UnregisterClient(userID uint, conn interface{})
	Broadcast(data interface{})
}
