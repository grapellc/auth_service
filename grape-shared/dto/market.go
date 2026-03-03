package dto

import (
	"github.com/your-moon/grape-shared/entities"
)

type MarketPricingDTO struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description"`
	Tag         string  `json:"tag"`
	Price       float64 `json:"price" binding:"required"`
	ImageID     *uint   `json:"image_id"`
}

type MarketCreateDTO struct {
	Name           string                    `json:"name"`
	Description    string                    `json:"description"`
	LocationID     *uint                     `json:"location_id"`
	Location       entities.LocationMapDetails `json:"location_details"`
	ImageID        *uint                     `json:"image_id"`
	ImageIDs       []uint                    `json:"image_ids"`
	MarketTypeID   *uint                     `json:"market_type_id"`
	OperatingHours *entities.OperatingHours    `json:"operating_hours"`
	Pricing        []MarketPricingDTO        `json:"pricing"`
	UserID         uint                      `json:"-"`
}

func (m *MarketCreateDTO) ToModel() *entities.Market {
	uid := m.UserID
	var pricing []entities.MarketPricing
	for _, p := range m.Pricing {
		pricing = append(pricing, entities.MarketPricing{
			Name:        p.Name,
			Description: p.Description,
			Tag:         p.Tag,
			Price:       p.Price,
			ImageID:     p.ImageID,
		})
	}

	market := &entities.Market{
		Name:               m.Name,
		Description:        m.Description,
		LocationID:         m.LocationID,
		LocationMapDetails: m.Location,
		ImageID:            m.ImageID,
		MarketTypeID:       m.MarketTypeID,
		OperatingHours:     m.OperatingHours,
		Pricing:            pricing,
	}
	market.CreatedByID = &uid
	market.UpdatedByID = &uid
	return market
}

type MarketUpdateDTO struct {
	Name           string                    `json:"name"`
	Description    string                    `json:"description"`
	LocationID     *uint                     `json:"location_id"`
	Location       entities.LocationMapDetails `json:"location_details"`
	ImageID        *uint                     `json:"image_id"`
	ImageIDs       []uint                    `json:"image_ids"`
	MarketTypeID   *uint                     `json:"market_type_id"`
	OperatingHours *entities.OperatingHours    `json:"operating_hours"`
	Pricing        []MarketPricingDTO        `json:"pricing"`
	UserID         uint                      `json:"-"`
}

func (m *MarketUpdateDTO) ToModel() *entities.Market {
	uid := m.UserID
	// Note: Update logic often differs (append vs replace). GORM creates new if ID missing.
	// For simple list replacement, we might need repository logic to clear old pricing.
	// But let's map it first.
	var pricing []entities.MarketPricing
	for _, p := range m.Pricing {
		pricing = append(pricing, entities.MarketPricing{
			Name:        p.Name,
			Description: p.Description,
			Tag:         p.Tag,
			Price:       p.Price,
			ImageID:     p.ImageID,
		})
	}

	market := &entities.Market{
		Name:               m.Name,
		Description:        m.Description,
		LocationID:         m.LocationID,
		LocationMapDetails: m.Location,
		ImageID:            m.ImageID,
		MarketTypeID:       m.MarketTypeID,
		OperatingHours:     m.OperatingHours,
		Pricing:            pricing,
	}
	market.UpdatedByID = &uid
	return market
}

// Review DTOs
type MarketReviewCreateDTO struct {
	MarketID uint   `json:"market_id"`
	Rating   int    `json:"rating" binding:"required,min=1,max=5"`
	Content  string `json:"content"`
	ImageIDs []uint `json:"image_ids"`
	UserID   uint   `json:"-"`
}

func (r *MarketReviewCreateDTO) ToModel() *entities.MarketReview {
	uid := r.UserID
	review := &entities.MarketReview{
		MarketID: r.MarketID,
		UserID:   r.UserID,
		Rating:   r.Rating,
		Content:  r.Content,
	}
	review.CreatedByID = &uid
	review.UpdatedByID = &uid
	return review
}

type MarketReviewUpdateDTO struct {
	Rating  *int    `json:"rating"`
	Content *string `json:"content"`
}

// Post DTOs
type MarketPostCreateDTO struct {
	MarketID uint   `json:"market_id"`
	Title    string `json:"title" binding:"required"`
	Content  string `json:"content"`
	ImageIDs []uint `json:"image_ids"`
	UserID   uint   `json:"-"`
}

func (p *MarketPostCreateDTO) ToModel() *entities.MarketPost {
	uid := p.UserID
	post := &entities.MarketPost{
		MarketID: p.MarketID,
		Title:    p.Title,
		Content:  p.Content,
	}
	post.CreatedByID = &uid
	post.UpdatedByID = &uid
	return post
}
