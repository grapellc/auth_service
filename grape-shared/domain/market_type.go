package domain

import (
	"context"
	"time"

	"github.com/your-moon/grape-shared/dto"
	"github.com/your-moon/grape-shared/common/types"
)

type MarketTypeService interface {
	Page(pagination types.Pagination) ([]*MarketType, int, error)
	GetByID(ctx context.Context, id uint) (*MarketType, error)
	Create(ctx context.Context, req dto.MarketTypeCreateDTO) (*MarketType, error)
	Update(ctx context.Context, id uint, req dto.MarketTypeUpdateDTO) (*MarketType, error)
	Delete(ctx context.Context, id uint) error
}

type MarketType struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
