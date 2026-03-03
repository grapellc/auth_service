package dto

type CreateConversationDTO struct {
	TargetUserID uint   `json:"target_user_id" validate:"required"`
	ProductID    *uint  `json:"product_id"` // Deprecated: Use ContextType/ID
	ContextType  string `json:"context_type"`
	ContextID    uint   `json:"context_id"`
}

type CreateMessageDTO struct {
	ConversationID uint   `json:"conversation_id" validate:"required"`
	Content        string `json:"content" validate:"required"`
}
