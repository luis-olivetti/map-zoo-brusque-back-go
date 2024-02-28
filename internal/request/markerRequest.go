package request

type MarkerRequest struct {
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Latitude    float64 `json:"lat" binding:"required"`
	Longitude   float64 `json:"lng" binding:"required"`
	Icon        string  `json:"icon" binding:"required"`
}
