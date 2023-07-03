package presentation

import "net/http"

type CalculateRideInput struct {
	Segments []Segment `json:"segments"`
}

type Segment struct {
	Distance float64 `json:"distance"`
	Date     string  `json:"date"`
}

type CalculateRideOutput struct {
	Price float64 `json:"price"`
}

func (o *CalculateRideOutput) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
