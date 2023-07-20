package entities

import "time"

type Position struct {
	Coord Coordinate
	Date  time.Time
}

func NewPosition(lat, long float64, date time.Time) *Position {
	coord := NewCoordinate(lat, long)
	return &Position{
		Coord: *coord,
		Date:  date,
	}
}
