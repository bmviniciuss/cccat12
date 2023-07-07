package presentation

import "net/http"

type CalculateRideInput struct {
	Segments []Segment `json:"segments"`
}
type Coordinate struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
}

type Segment struct {
	From Coordinate `json:"from"`
	To   Coordinate `json:"to"`
	Date string     `json:"date"`
}

type CalculateRideOutput struct {
	Price float64 `json:"price"`
}

func (o *CalculateRideOutput) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
