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
	Positions   []Position
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
		Positions: []Position{},
	}
}

func (r *Ride) AddPosition(lat, long float64, date time.Time) {
	pos := NewPosition(lat, long, date)
	r.Positions = append(r.Positions, *pos)
}

func (r *Ride) Calculate() (float64, error) {
	price := 0.0
	for i := 0; i < len(r.Positions)-1; i++ {
		current := r.Positions[i]
		next := r.Positions[i+1]
		distance := current.Coord.DistanceInMeters(next.Coord)
		segment, err := NewSegment(distance, next.Date)
		if err != nil {
			return -1, err
		}
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
		return MinPrice, nil
	}
	return price, nil
}
