package entities

import (
	"time"
)

const (
	OvernightFare       float64 = 3.90
	OvernightSundayFare float64 = 5.0
	SundayFare          float64 = 2.9
	NormalFare          float64 = 2.1
	MinPrice            float64 = 10
)

type Ride struct {
	ID          string
	PassengerID string
	DriverID    *string
	RequestDate time.Time
	From        Coordinate
	To          Coordinate
	Segments    []Segment
}

func CreateRide(passengerID string, from, to Coordinate) *Ride {
	return &Ride{
		ID:          NewULID().String(),
		PassengerID: passengerID,
		From:        from,
		To:          to,
		RequestDate: time.Now(),
		DriverID:    nil,
	}
}

func NewRide() *Ride {
	return &Ride{
		Segments: []Segment{},
	}
}

func (r *Ride) AddSegment(from, to Coordinate, date time.Time) error {
	seg, err := NewSegment(from, to, date)
	if err != nil {
		return err
	}
	r.Segments = append(r.Segments, *seg)
	return nil
}

func (r *Ride) Calculate() float64 {
	price := 0.0
	for _, segment := range r.Segments {
		if segment.IsOvernight() && !segment.IsSunday() {
			price += segment.Distance * OvernightFare
		}
		if segment.IsOvernight() && segment.IsSunday() {
			price += segment.Distance * OvernightSundayFare
		}
		if !segment.IsOvernight() && segment.IsSunday() {
			price += segment.Distance * SundayFare
		}
		if !segment.IsOvernight() && !segment.IsSunday() {
			price += segment.Distance * NormalFare
		}
	}

	if price < MinPrice {
		return MinPrice
	}

	return price
}
