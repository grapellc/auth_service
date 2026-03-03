package dmo

import (
	"github.com/your-moon/grape-shared/bc/jobs"
	"github.com/your-moon/grape-shared/bc/market"
	"github.com/your-moon/grape-shared/bc/product"
	"github.com/your-moon/grape-shared/domain"
)

// ToSharedProduct converts product bounded-context domain to shared domain.
func ToSharedProduct(p *product.Product) *domain.Product {
	if p == nil {
		return nil
	}
	var cat *domain.Category
	if p.Category != nil {
		cat = &domain.Category{ID: p.Category.ID, Name: p.Category.Name}
	}
	return &domain.Product{
		ID:             p.ID,
		Name:           p.Name,
		Description:    p.Description,
		Price:          p.Price,
		ViewCount:      p.ViewCount,
		Status:         p.Status,
		Category:       cat,
		Location:       p.Location,
		LocationDetail: p.LocationDetail,
		Images:         p.Images,
		Thumbnail:      p.Thumbnail,
		CreatedAt:      p.CreatedAt,
		UpdatedAt:      p.UpdatedAt,
		CreatedByID:    p.CreatedByID,
		CreatedBy:      p.CreatedBy,
		UpdatedByID:    p.UpdatedByID,
		IsLiked:        p.IsLiked,
		LikeCount:      p.LikeCount,
	}
}

// ToSharedJob converts jobs bounded-context domain to shared domain.
func ToSharedJob(j *jobs.Job) *domain.Job {
	if j == nil {
		return nil
	}
	var companyInfo *domain.JobCompanyInfo
	if j.CompanyInfo != nil {
		companyInfo = &domain.JobCompanyInfo{
			ID:          j.CompanyInfo.ID,
			CompanyName: j.CompanyInfo.CompanyName,
			Email:       j.CompanyInfo.Email,
			Phone:       j.CompanyInfo.Phone,
			Salary:      j.CompanyInfo.Salary,
		}
	}
	return &domain.Job{
		ID:             j.ID,
		Name:           j.Name,
		Description:    j.Description,
		JobType:        j.JobType,
		HourlyRate:     j.HourlyRate,
		StartDate:      j.StartDate,
		EndDate:        j.EndDate,
		Status:         j.Status,
		ViewCount:      j.ViewCount,
		Location:       j.Location,
		LocationDetail: j.LocationDetail,
		CompanyInfo:    companyInfo,
		Images:         j.Images,
		Thumbnail:      j.Thumbnail,
		CreatedAt:      j.CreatedAt,
		UpdatedAt:      j.UpdatedAt,
		CreatedByID:    j.CreatedByID,
		CreatedBy:      j.CreatedBy,
		UpdatedByID:    j.UpdatedByID,
		IsLiked:        j.IsLiked,
		LikeCount:      j.LikeCount,
	}
}

// ToSharedMarket converts market bounded-context domain to shared domain.
func ToSharedMarket(m *market.Market) *domain.Market {
	if m == nil {
		return nil
	}
	var loc *domain.Location
	if m.Location != nil {
		loc = toSharedLocationFromMarket(m.Location)
	}
	var user *domain.User
	if m.CreatedBy != nil {
		user = toSharedUserFromMarket(m.CreatedBy)
	}
	var mt *domain.MarketType
	if m.MarketType != nil {
		mt = &domain.MarketType{
			ID:          m.MarketType.ID,
			Name:        m.MarketType.Name,
			Description: m.MarketType.Description,
			CreatedAt:   m.MarketType.CreatedAt,
			UpdatedAt:   m.MarketType.UpdatedAt,
		}
	}
	var reviews []*domain.MarketReview
	for _, r := range m.Reviews {
		reviews = append(reviews, toSharedMarketReview(r))
	}
	var posts []*domain.MarketPost
	for _, p := range m.Posts {
		posts = append(posts, toSharedMarketPost(p))
	}
	var pricing []*domain.MarketPricing
	for _, p := range m.Pricing {
		pricing = append(pricing, &domain.MarketPricing{
			ID:          p.ID,
			MarketID:    p.MarketID,
			Name:        p.Name,
			Description: p.Description,
			Tag:         p.Tag,
			Price:       p.Price,
			Image:       p.Image,
		})
	}
	return &domain.Market{
		ID:             m.ID,
		Name:           m.Name,
		Description:    m.Description,
		ViewCount:      m.ViewCount,
		Image:          m.Image,
		Images:         m.Images,
		Location:       loc,
		LocationDetail: m.LocationDetail,
		MarketType:     mt,
		OperatingHours: m.OperatingHours,
		CreatedAt:      m.CreatedAt,
		UpdatedAt:      m.UpdatedAt,
		CreatedByID:    m.CreatedByID,
		CreatedBy:      user,
		UpdatedByID:    m.UpdatedByID,
		IsLiked:        m.IsLiked,
		LikeCount:      m.LikeCount,
		Reviews:        reviews,
		Posts:          posts,
		Pricing:        pricing,
		AverageRating:  m.AverageRating,
		ReviewCount:    m.ReviewCount,
	}
}

func toSharedLocationFromMarket(l *market.Location) *domain.Location {
	if l == nil {
		return nil
	}
	var area *domain.LocationArea
	if l.LocationArea != nil {
		area = &domain.LocationArea{ID: l.LocationArea.ID, Name: l.LocationArea.Name}
	}
	return &domain.Location{
		ID:           l.ID,
		Name:         l.Name,
		Address:      l.Address,
		Lat:          l.Lat,
		Lng:          l.Lng,
		LocationArea: area,
	}
}

func toSharedUserFromMarket(u *market.User) *domain.User {
	if u == nil {
		return nil
	}
	return &domain.User{
		ID:              u.ID,
		Name:            u.Name,
		Email:           u.Email,
		Username:        u.Username,
		AvatarURL:       u.AvatarURL,
		CreatedAt:       u.CreatedAt,
		UpdatedAt:       u.UpdatedAt,
		IsPhoneVerified: u.IsPhoneVerified,
		IsEmailVerified: u.IsEmailVerified,
	}
}

func toSharedMarketReview(r *market.MarketReview) *domain.MarketReview {
	if r == nil {
		return nil
	}
	var u *domain.User
	if r.User != nil {
		u = toSharedUserFromMarket(r.User)
	}
	return &domain.MarketReview{
		ID:         r.ID,
		MarketID:   r.MarketID,
		UserID:     r.UserID,
		Rating:     r.Rating,
		Content:    r.Content,
		LikesCount: r.LikesCount,
		IsLiked:    r.IsLiked,
		Images:     r.Images,
		User:       u,
		CreatedAt:  r.CreatedAt,
		UpdatedAt:  r.UpdatedAt,
	}
}

func toSharedMarketPost(p *market.MarketPost) *domain.MarketPost {
	if p == nil {
		return nil
	}
	return &domain.MarketPost{
		ID:           p.ID,
		MarketID:     p.MarketID,
		Title:        p.Title,
		Content:      p.Content,
		LikeCount:    p.LikeCount,
		CommentCount: p.CommentCount,
		Images:       p.Images,
		CreatedAt:    p.CreatedAt,
		UpdatedAt:    p.UpdatedAt,
	}
}
