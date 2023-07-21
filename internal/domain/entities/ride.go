package entities

import (
	"time"
)

const (
	minPrice = 10.00
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
		fareCalculator, err := createFareCalculator(*segment)
		if err != nil {
			return -1, err
		}

		price += fareCalculator.Calculate(*segment)
	}
	if price < minPrice {
		return minPrice, nil
	}
	return price, nil
}
