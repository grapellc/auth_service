package dto

import "github.com/your-moon/grape-shared/entities"

type ProductCreateDTO struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Price       float64                   `json:"price"`
	UserID      uint                      `json:"user_id"`
	CategoryID  *uint                     `json:"category_id,omitempty"`
	LocationID  *uint                     `json:"location_id,omitempty"`
	Location    entities.LocationMapDetails `json:"location_details"`
	ImageIDs    []uint                    `json:"image_ids,omitempty"`
}

func (p *ProductCreateDTO) ToModel() *entities.Product {
	uid := p.UserID
	product := &entities.Product{
		Name:               p.Name,
		Description:        p.Description,
		Price:              p.Price,
		LocationMapDetails: p.Location,
	}
	product.CreatedByID = &uid
	product.UpdatedByID = &uid

	// Only set CategoryID if provided and not zero
	if p.CategoryID != nil && *p.CategoryID > 0 {
		product.CategoryID = p.CategoryID
	}

	// Only set LocationID if provided and not zero
	if p.LocationID != nil && *p.LocationID > 0 {
		product.LocationID = p.LocationID
	}

	return product
}

type ProductUpdateDTO struct {
	Name        string                    `json:"name"`
	Description string                    `json:"description"`
	Price       float64                   `json:"price"`
	CategoryID  *uint                     `json:"category_id,omitempty"`
	LocationID  *uint                     `json:"location_id,omitempty"`
	Location    entities.LocationMapDetails `json:"location_details"`
	ImageIDs    []uint                    `json:"image_ids,omitempty"`
}

func (p *ProductUpdateDTO) ToModel() *entities.Product {
	product := &entities.Product{
		Name:               p.Name,
		Description:        p.Description,
		Price:              p.Price,
		LocationMapDetails: p.Location,
	}

	// Only set CategoryID if provided and not zero
	if p.CategoryID != nil && *p.CategoryID > 0 {
		product.CategoryID = p.CategoryID
	}

	// Only set LocationID if provided and not zero
	if p.LocationID != nil && *p.LocationID > 0 {
		product.LocationID = p.LocationID
	}

	return product
}
