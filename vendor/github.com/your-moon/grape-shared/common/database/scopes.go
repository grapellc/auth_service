package database

import "gorm.io/gorm"

func Search(idList []uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(idList) > 0 {
			return db.Where("id IN ?", idList)
		}
		return db
	}
}

func FilterByLocation(locationID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if locationID != 0 {
			return db.Where("location_id = ?", locationID)
		}
		return db
	}
}

func FilterByCategory(categoryID uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if categoryID != 0 {
			return db.Where("category_id = ?", categoryID)
		}
		return db
	}
}

func FilterByCategories(categoryIDs []uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		// Filter out zero values (empty category_id parameters)
		filtered := make([]uint, 0)
		for _, id := range categoryIDs {
			if id != 0 {
				filtered = append(filtered, id)
			}
		}
		if len(filtered) > 0 {
			return db.Where("category_id IN ?", filtered)
		}
		return db
	}
}

func FilterByPriceRange(priceMin, priceMax float64) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if priceMin > 0 {
			db = db.Where("price >= ?", priceMin)
		}
		if priceMax > 0 {
			db = db.Where("price <= ?", priceMax)
		}
		return db
	}
}

func FilterByJobType(jobType string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if jobType != "" && jobType != "all" {
			return db.Where("job_type = ?", jobType)
		}
		return db
	}
}

func FilterByJobTypes(jobTypes []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(jobTypes) > 0 {
			// Filter out "all" values
			filtered := make([]string, 0)
			for _, jt := range jobTypes {
				if jt != "" && jt != "all" {
					filtered = append(filtered, jt)
				}
			}
			if len(filtered) > 0 {
				return db.Where("job_type IN ?", filtered)
			}
		}
		return db
	}
}

func FilterByStatus(status string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if status != "" && status != "all" {
			return db.Where("status = ?", status)
		}
		return db
	}
}

func FilterByStatuses(statuses []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(statuses) > 0 {
			// Filter out "all" values
			filtered := make([]string, 0)
			for _, st := range statuses {
				if st != "" && st != "all" {
					filtered = append(filtered, st)
				}
			}
			if len(filtered) > 0 {
				return db.Where("status IN ?", filtered)
			}
		}
		return db
	}
}

func FilterByMarketTypes(marketTypeIDs []uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if len(marketTypeIDs) > 0 {
			// Filter out zero values
			filtered := make([]uint, 0)
			for _, id := range marketTypeIDs {
				if id != 0 {
					filtered = append(filtered, id)
				}
			}
			if len(filtered) > 0 {
				return db.Where("market_type_id IN ?", filtered)
			}
		}
		return db
	}
}

func FilterByCreator(userID *uint) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if userID != nil && *userID > 0 {
			return db.Where("created_by_id = ?", *userID)
		}
		return db
	}
}
