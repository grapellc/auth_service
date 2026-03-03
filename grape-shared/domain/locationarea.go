package domain

import (
	"context"

	"github.com/your-moon/grape-shared/entities"
	"github.com/your-moon/grape-shared/common/types"
)

type LocationAreaService interface {
	ListLocationAreas(ctx context.Context, pagination types.Pagination) ([]*entities.LocationArea, int, error)
	GetLocationArea(ctx context.Context, id uint) (*entities.LocationArea, error)
	CreateLocationArea(ctx context.Context, locationArea *entities.LocationArea) (*entities.LocationArea, error)
	UpdateLocationArea(ctx context.Context, id uint, locationArea *entities.LocationArea) (*entities.LocationArea, error)
	DeleteLocationArea(ctx context.Context, id uint) error
}

type LocationArea struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}
