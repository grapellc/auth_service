package domain

import (
	"context"

	"github.com/your-moon/grape-shared/common/types"
)

type LocationService interface {
	ListLocations(ctx context.Context, pagination types.Pagination) ([]*Location, int, error)
	GetLocation(ctx context.Context, id uint) (*Location, error)
	// Create and Update still accept *entities.Location for input simplicity for now,
	// but typically should accept DTOs. Let's keep input as models for now to minimize disruption,
	// or better yet, switch to DTOs. The Controller creates models from body parser.
	// For strict DDD, we should accept a CreateLocationDTO.
	// However, to match the current refactoring scope (Output standardization), let's stick to returning Domain entities.
	// But wait, the interface signature defined in previous step uses *entities.Location for inputs.
	// Let's keep inputs as is for this specific task to avoid cascading changes to controllers,
	// but fix the return types.
	// ACTUALLY, looking at other services (Job), Create takes a DTO.
	// LocationController uses `entities.Location` for body parsing.
	// Let's fix the Return types first.
	CreateLocation(ctx context.Context, location *Location) (*Location, error)
	UpdateLocation(ctx context.Context, id uint, location Location) (*Location, error)
	DeleteLocation(ctx context.Context, id uint) error
}

type Location struct {
	ID           uint          `json:"id"`
	Name         string        `json:"name"`
	Address      string        `json:"address"`
	Lat          float64       `json:"lat"`
	Lng          float64       `json:"lng"`
	LocationArea *LocationArea `json:"location_area,omitempty"`
}
