package dmo

import (
	"github.com/your-moon/grape-shared/domain"
	"github.com/your-moon/grape-shared/entities"
)

func ToProductDomain(product *entities.Product, isLiked bool, likeCount int) *domain.Product {
	// Convert images to string URLs
	var imageURLs []string
	var thumbnail string
	if product.Images != nil {
		for i, img := range product.Images {
			if img.File != nil && img.File.URL != nil {
				imageURLs = append(imageURLs, *img.File.URL)
				// Set thumbnail to the first image
				if i == 0 {
					thumbnail = *img.File.URL
				}
			}
		}
	}

	// Convert user
	var user *domain.User
	if product.CreatedBy != nil {
		// Combine first and last name for display name
		var displayName string
		if product.CreatedBy.FirstName != nil && product.CreatedBy.LastName != nil {
			displayName = *product.CreatedBy.FirstName + " " + *product.CreatedBy.LastName
		} else if product.CreatedBy.FirstName != nil {
			displayName = *product.CreatedBy.FirstName
		} else if product.CreatedBy.LastName != nil {
			displayName = *product.CreatedBy.LastName
		} else if product.CreatedBy.Username != nil {
			displayName = *product.CreatedBy.Username
		}

		user = &domain.User{
			ID:    product.CreatedBy.ID,
			Name:  displayName,
			Email: product.CreatedBy.Email,
		}
	}

	// Convert category
	var category *domain.Category
	if product.Category != nil {
		category = &domain.Category{
			ID:   product.Category.ID,
			Name: product.Category.Name,
		}
	}

	// Convert location
	var location *domain.Location
	if product.Location != nil {
		location = ToLocation(product.Location)
	} else if product.Latitude != nil && product.Longitude != nil {
		name := "Сонгосон байршил"
		if product.Locality != nil && *product.Locality != "" {
			name = *product.Locality
		} else if product.Region != nil && *product.Region != "" {
			name = *product.Region
		}

		location = &domain.Location{
			Name: name,
			Lat:  *product.Latitude,
			Lng:  *product.Longitude,
		}
	}

	locationDetail := entities.LocationMapDetails{
		Latitude:  product.Latitude,
		Longitude: product.Longitude,
		Postcode:  product.Postcode,
		Locality:  product.Locality,
		Region:    product.Region,
		Country:   product.Country,
	}

	return &domain.Product{
		ID:             product.ID,
		Name:           product.Name,
		Description:    product.Description,
		Price:          product.Price,
		ViewCount:      product.ViewCount,
		Status:         product.Status,
		LocationDetail: locationDetail,
		CreatedBy:      user,
		Category:       category,
		Location:       location,
		Images:         imageURLs,
		Thumbnail:      thumbnail,
		CreatedAt:      product.CreatedAt,
		UpdatedAt:      product.UpdatedAt,
		CreatedByID:    product.CreatedByID,
		UpdatedByID:    product.UpdatedByID,
		IsLiked:        isLiked,
		LikeCount:      int64(likeCount),
	}
}
