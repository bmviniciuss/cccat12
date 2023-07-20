package presentation

import "net/http"

type CalculateRideInput struct {
	Positions []Position `json:"positions"`
}

type Position struct {
	Lat  float64 `json:"lat"`
	Long float64 `json:"long"`
	Date string  `json:"date"`
}

type CalculateRideOutput struct {
	Price float64 `json:"price"`
}

func (o *CalculateRideOutput) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
