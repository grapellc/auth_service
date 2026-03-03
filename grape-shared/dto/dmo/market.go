package dmo

import (
	"github.com/your-moon/grape-shared/domain"
	"github.com/your-moon/grape-shared/entities"
)

func ToDomain(market *entities.Market, isLiked bool, likeCount int64) *domain.Market {
	if market == nil {
		return nil
	}

	// Convert location
	var location *domain.Location
	if market.Location != nil {
		location = ToLocation(market.Location)
	} else if market.Latitude != nil && market.Longitude != nil {
		// Fallback to coordinates if Location relation is missing
		name := "Сонгосон байршил"
		if market.Locality != nil && *market.Locality != "" {
			name = *market.Locality
		} else if market.Region != nil && *market.Region != "" {
			name = *market.Region
		}

		location = &domain.Location{
			Name: name,
			Lat:  *market.Latitude,
			Lng:  *market.Longitude,
		}
	}

	// Convert user
	var user *domain.User
	if market.Creator != nil {
		user = ToUserDomain(market.Creator)
	} else if market.CreatedBy != nil {
		user = ToUserDomain(market.CreatedBy)
	}

	// Convert images
	var images []string
	if len(market.Images) > 0 {
		for _, img := range market.Images {
			if img.File != nil && img.File.URL != nil {
				images = append(images, *img.File.URL)
			}
		}
	}
	// Fallback to single image if no multiple images
	if len(images) == 0 && market.Image != nil && market.Image.URL != nil {
		images = []string{*market.Image.URL}
	}

	// Convert reviews
	var reviews []*domain.MarketReview
	if len(market.Reviews) > 0 {
		for _, review := range market.Reviews {
			reviews = append(reviews, ToMarketReviewDomain(review, false)) // Defaults to false in nested view
		}
	}

	// Convert posts
	var posts []*domain.MarketPost
	if len(market.Posts) > 0 {
		for _, post := range market.Posts {
			posts = append(posts, ToMarketPostDomain(post))
		}
	}

	return &domain.Market{
		ID:             market.ID,
		Name:           market.Name,
		Description:    market.Description,
		ViewCount:      market.ViewCount,
		Location:       location,
		LocationDetail: market.LocationMapDetails,
		MarketType:     ToMarketTypeDomain(market.MarketType),
		OperatingHours: market.OperatingHours,
		Image: func() string {
			if len(images) > 0 {
				return images[0]
			}
			if market.Image != nil && market.Image.URL != nil {
				return *market.Image.URL
			}
			return ""
		}(),
		Images:        images,
		CreatedAt:     market.CreatedAt,
		UpdatedAt:     market.UpdatedAt,
		CreatedByID:   market.CreatedByID,
		CreatedBy:     user,
		UpdatedByID:   market.UpdatedByID,
		IsLiked:       isLiked,
		LikeCount:     likeCount,
		Reviews:       reviews,
		Posts:         posts,
		AverageRating: market.AverageRating,
		ReviewCount:   market.ReviewCount,
		Pricing: func() []*domain.MarketPricing {
			var pricing []*domain.MarketPricing
			if len(market.Pricing) > 0 {
				for _, p := range market.Pricing {
					pricing = append(pricing, ToMarketPricingDomain(&p))
				}
			}
			return pricing
		}(),
	}
}

func ToMarketPricingDomain(pricing *entities.MarketPricing) *domain.MarketPricing {
	if pricing == nil {
		return nil
	}
	image := ""
	if pricing.Image != nil && pricing.Image.URL != nil {
		image = *pricing.Image.URL
	}
	return &domain.MarketPricing{
		ID:          pricing.ID,
		MarketID:    pricing.MarketID,
		Name:        pricing.Name,
		Description: pricing.Description,
		Tag:         pricing.Tag,
		Price:       pricing.Price,
		Image:       image,
	}
}

func ToMarketTypeDomain(marketType *entities.MarketType) *domain.MarketType {
	if marketType == nil {
		return nil
	}
	return &domain.MarketType{
		ID:          marketType.ID,
		Name:        marketType.Name,
		Description: marketType.Description,
		CreatedAt:   marketType.CreatedAt,
		UpdatedAt:   marketType.UpdatedAt,
	}
}

func ToMarketReviewDomain(review *entities.MarketReview, isLiked bool) *domain.MarketReview {
	if review == nil {
		return nil
	}

	var images []string
	if len(review.Images) > 0 {
		for _, img := range review.Images {
			if img.File != nil && img.File.URL != nil {
				images = append(images, *img.File.URL)
			}
		}
	}

	return &domain.MarketReview{
		ID:         review.ID,
		MarketID:   review.MarketID,
		UserID:     review.UserID,
		Rating:     review.Rating,
		Content:    review.Content,
		LikesCount: review.LikesCount,
		IsLiked:    isLiked,
		Images:     images,
		User:       ToUserDomain(review.User),
		CreatedAt:  review.CreatedAt,
		UpdatedAt:  review.UpdatedAt,
	}
}

func ToMarketPostDomain(post *entities.MarketPost) *domain.MarketPost {
	if post == nil {
		return nil
	}

	var images []string
	if len(post.Images) > 0 {
		for _, img := range post.Images {
			if img.File != nil && img.File.URL != nil {
				images = append(images, *img.File.URL)
			}
		}
	}

	return &domain.MarketPost{
		ID:           post.ID,
		MarketID:     post.MarketID,
		Title:        post.Title,
		Content:      post.Content,
		LikeCount:    post.LikeCount,
		CommentCount: post.CommentCount,
		Images:       images,
		CreatedAt:    post.CreatedAt,
		UpdatedAt:    post.UpdatedAt,
	}
}
