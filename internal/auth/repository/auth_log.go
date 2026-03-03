package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/your-moon/grape-shared/entities"
	"gorm.io/gorm"
)

type authLogRepository struct {
	db *gorm.DB
}

func NewAuthLogRepository(db *gorm.DB) *authLogRepository {
	return &authLogRepository{db: db}
}

func (r *authLogRepository) Create(ctx context.Context, log *entities.AuthLogUUID) error {
	if log.ID == uuid.Nil {
		log.ID = uuid.New()
	}
	return r.db.WithContext(ctx).Create(log).Error
}
