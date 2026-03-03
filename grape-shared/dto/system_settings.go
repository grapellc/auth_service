package dto

type SystemSettings struct {
	MapStyle      string            `json:"mapStyle"`
	Modules       map[string]string `json:"modules"`
	MobileModules map[string]string `json:"mobile_modules"`
}
