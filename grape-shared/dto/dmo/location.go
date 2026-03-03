package dmo

import (
	"github.com/your-moon/grape-shared/domain"
	"github.com/your-moon/grape-shared/entities"
)

func ToLocation(location *entities.Location) *domain.Location {
	if location == nil {
		return nil
	}
	return &domain.Location{
		ID:           location.ID,
		Name:         location.Name,
		Address:      location.Address,
		Lat:          location.Latitude,
		Lng:          location.Longitude,
		LocationArea: ToLocationArea(location.LocationArea),
	}
}

func ToLocationModel(location *domain.Location) *entities.Location {
	if location == nil {
		return nil
	}
	return &entities.Location{
		BaseModel: entities.BaseModel{ID: location.ID},
		Name:      location.Name,
		Address:   location.Address,
		Latitude:  location.Lat,
		Longitude: location.Lng,
	}
}

func ToLocationArea(locationArea *entities.LocationArea) *domain.LocationArea {
	if locationArea == nil {
		return nil
	}
	return &domain.LocationArea{
		ID:   locationArea.ID,
		Name: locationArea.Name,
	}
}
