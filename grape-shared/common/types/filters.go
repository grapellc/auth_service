package types

// ProductFilters contains product-specific filter fields
type ProductFilters struct {
	CategoryIDs []uint  `query:"category_id"`
	PriceMin    float64 `query:"price_min"`
	PriceMax    float64 `query:"price_max"`
	Sort        string  `query:"sort" default:"created_at DESC"`
	UserID      *uint   `query:"user_id"`
}

// JobFilters contains job-specific filter fields
type JobFilters struct {
	JobTypes []string `query:"job_type[]"`
	Statuses []string `query:"status[]"`
}

// ProductPagination embeds Pagination and ProductFilters
type ProductPagination struct {
	Pagination
	ProductFilters
	CurrentUserID *uint `json:"-" query:"-"` // Not bound from query params, set by service
}

// JobPagination embeds Pagination and JobFilters
type JobPagination struct {
	Pagination
	JobFilters
}

// GetSort returns the SQL ORDER BY clause for product sorting
func (pf *ProductFilters) GetSort() string {
	if pf.Sort == "" {
		return DEFAULT_SORTER
	}
	// Map frontend sort values to SQL ORDER BY
	sortMap := map[string]string{
		"newest":     "created_at DESC",
		"oldest":     "created_at ASC",
		"price_asc":  "price ASC",
		"price_desc": "price DESC",
		"name_asc":   "name ASC",
		"name_desc":  "name DESC",
	}
	if mapped, ok := sortMap[pf.Sort]; ok {
		return mapped
	}
	return pf.Sort
}
