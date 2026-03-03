package dmo

import (
	"github.com/your-moon/grape-shared/domain"
	"github.com/your-moon/grape-shared/entities"
)

func ToUserDomain(user *entities.User) *domain.User {
	if user == nil {
		return nil
	}

	domainUser := &domain.User{
		ID:              user.ID,
		Email:           user.Email,
		CreatedAt:       user.CreatedAt,
		UpdatedAt:       user.UpdatedAt,
		CreatedUserID:   user.CreatedUserID,
		UpdatedUserID:   user.UpdatedUserID,
		AvatarURL:       user.AvatarURL,
		IsPhoneVerified: user.IsPhoneVerified,
		IsEmailVerified: user.IsEmailVerified,
	}

	if user.Username != nil {
		domainUser.Username = user.Username
	}
	if user.FirstName != nil {
		domainUser.FirstName = user.FirstName
	}
	if user.LastName != nil {
		domainUser.LastName = user.LastName
	}
	if user.PhoneNumber != nil {
		domainUser.PhoneNumber = user.PhoneNumber
	}

	// Calculate Display Name
	if user.FirstName != nil && user.LastName != nil {
		domainUser.Name = *user.FirstName + " " + *user.LastName
	} else if user.FirstName != nil {
		domainUser.Name = *user.FirstName
	} else if user.LastName != nil {
		domainUser.Name = *user.LastName
	} else if user.Username != nil {
		domainUser.Name = *user.Username
	}

	return domainUser
}
