package presentation

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
