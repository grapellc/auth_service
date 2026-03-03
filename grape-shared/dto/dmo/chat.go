package dmo

import (
	"github.com/your-moon/grape-shared/domain"
	"github.com/your-moon/grape-shared/entities"
)

func ToDomainConversation(m *entities.Conversation, currentUserID uint, unreadCount int) domain.Conversation {
	var otherUser *entities.User
	if m.User1ID == currentUserID {
		otherUser = &m.User2
	} else {
		otherUser = &m.User1
	}

	var context *domain.ConversationContext

	// Handle Legacy Product Logic & Generic Context
	if m.ContextType == "product" || (m.Product != nil && m.ContextType == "") {
		// Fallback for when ContextType isn't explicitly set but ProductID is
		if m.Product != nil {
			imageUrl := ""
			if len(m.Product.Images) > 0 && m.Product.Images[0].File != nil && m.Product.Images[0].File.URL != nil {
				imageUrl = *m.Product.Images[0].File.URL
			}
			context = &domain.ConversationContext{
				ID:       m.Product.ID,
				Type:     "product",
				Title:    m.Product.Name,
				Subtitle: "", // Could be price formatted
				ImageURL: imageUrl,
			}
		}
	} else if m.ContextType == "job" && m.Job != nil {
		context = &domain.ConversationContext{
			ID:       m.Job.ID,
			Type:     "job",
			Title:    m.Job.Name,
			Subtitle: "",                                      // Could be company name or salary
			ImageURL: "https://placehold.co/100x100?text=JOB", // Job might not have an image
		}
	} else if m.ContextType == "market" && m.Market != nil {
		imageUrl := ""
		if m.Market.Image != nil && m.Market.Image.URL != nil {
			imageUrl = *m.Market.Image.URL
		}
		context = &domain.ConversationContext{
			ID:       m.Market.ID,
			Type:     "market",
			Title:    m.Market.Name,
			Subtitle: "", // Could be price
			ImageURL: imageUrl,
		}
	}

	var isPinned bool
	if m.User1ID == currentUserID {
		isPinned = m.PinnedByUser1
	} else {
		isPinned = m.PinnedByUser2
	}

	return domain.Conversation{
		ID:            m.ID,
		User:          ToUserDomain(otherUser),
		LastMessage:   m.LastMessage,
		LastMessageAt: m.LastMessageAt,
		UnreadCount:   unreadCount,
		IsBuying:      m.User1ID == currentUserID,
		IsSelling:     m.User2ID == currentUserID,
		IsPinned:      isPinned,
		Context:       context,
	}
}

func ToDomainMessage(m *entities.Message) domain.Message {
	return domain.Message{
		ID:             m.ID,
		ConversationID: m.ConversationID,
		SenderID:       m.SenderID,
		Sender:         ToUserDomain(&m.Sender),
		Content:        m.Content,
		CreatedAt:      m.CreatedAt,
		IsRead:         m.IsRead,
	}
}
