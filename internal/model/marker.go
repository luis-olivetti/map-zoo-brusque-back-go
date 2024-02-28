package model

type Marker struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Latitude    float64 `json:"lat"`
	Longitude   float64 `json:"lng"`
	Icon        string  `json:"icon"`
}
