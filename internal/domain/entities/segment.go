package entities

import (
	"errors"
	"time"
)

var (
	ErrInvalidSegmentDistance = errors.New("invalid segment distance")
)

var TimeLayout = "2006-01-02T15:04:05"

type Segment struct {
	Distance float64
	From     Coordinate
	To       Coordinate
	Date     time.Time
}

func NewSegment(from, to Coordinate, date time.Time) (*Segment, error) {
	seg := &Segment{
		Distance: from.DistanceInMeters(to),
		From:     from,
		To:       to,
		Date:     date,
	}
	err := seg.valid()
	if err != nil {
		return nil, err
	}
	return seg, nil
}

func (s Segment) valid() error {
	if s.Distance <= 0 {
		return ErrInvalidSegmentDistance
	}
	return nil
}

func (s Segment) IsOvernight() bool {
	return s.Date.Hour() >= 22 || s.Date.Hour() <= 6
}

func (s Segment) IsSunday() bool {
	return s.Date.Weekday() == time.Sunday
}
