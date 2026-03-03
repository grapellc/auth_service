package constants

const (
	// Product validation constants
	MinProductNameLength        = 1
	MaxProductNameLength        = 255
	MaxProductDescriptionLength = 1000
	MinProductPrice             = 0
	MaxProductPrice             = 999999999

	// Product status
	ProductStatusActive   = "active"
	ProductStatusInactive = "inactive"
	ProductStatusSold     = "sold"

	// Pagination
	DefaultPageSize = 20
	MaxPageSize     = 100

	ProductIndexName = "products"
)
