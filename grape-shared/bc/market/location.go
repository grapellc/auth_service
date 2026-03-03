package market

type LocationArea struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Location struct {
	ID           uint          `json:"id"`
	Name         string        `json:"name"`
	Address      string        `json:"address"`
	Lat          float64       `json:"lat"`
	Lng          float64       `json:"lng"`
	LocationArea *LocationArea `json:"location_area,omitempty"`
}
