package dto

import "github.com/your-moon/grape-shared/entities"

type MarketTypeCreateDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"-"`
}

func (m *MarketTypeCreateDTO) ToModel() *entities.MarketType {
	uid := m.UserID
	mt := &entities.MarketType{
		Name:        m.Name,
		Description: m.Description,
	}
	mt.CreatedByID = &uid
	mt.UpdatedByID = &uid
	return mt
}

type MarketTypeUpdateDTO struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      uint   `json:"-"`
}

func (m *MarketTypeUpdateDTO) ToModel() *entities.MarketType {
	uid := m.UserID
	mt := &entities.MarketType{
		Name:        m.Name,
		Description: m.Description,
	}
	mt.UpdatedByID = &uid
	return mt
}
