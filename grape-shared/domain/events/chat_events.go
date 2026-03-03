package events

type MessageCreatedEvent struct {
	MessageID      uint   `json:"message_id"`
	ConversationID uint   `json:"conversation_id"`
	SenderID       uint   `json:"sender_id"`
	Content        string `json:"content"`
	RecipientID    uint   `json:"recipient_id"`
}
