package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewRide(t *testing.T) {
	t.Run("should return a new ride with empty segments", func(t *testing.T) {
		ride := NewRide()
		assert.Equal(t, ride.Segments, []Segment{})
		assert.Equal(t, len(ride.Segments), 0)
	})
}

var layout = "2006-01-02T15:04:05"

func TestRide_Calculate(t *testing.T) {
	t.Run("should return an ErrInvalidSegmentDistance when segment distance is invalid", func(t *testing.T) {
		ride := NewRide()
		from := *NewCoordinate(0, 0)
		to := *NewCoordinate(0, 0)
		err := ride.AddSegment(from, to, time.Now())
		assert.Equal(t, err, ErrInvalidSegmentDistance)
	})

	t.Run("should calcule a ride during day time", func(t *testing.T) {
		ride := NewRide()
		time, _ := time.Parse(layout, "2021-03-01T10:00:00")
		from := *NewCoordinate(40.7177, -74.0060)
		to := *NewCoordinate(40.6385, -74.0060)
		err := ride.AddSegment(from, to, time)
		price := ride.Calculate()
		assert.Nil(t, err)
		assert.Equal(t, 18493.94019952293, price)
	})

	t.Run("should calcule a ride during night time", func(t *testing.T) {
		ride := NewRide()
		time, _ := time.Parse(layout, "2021-03-01T23:00:00")
		from := *NewCoordinate(40.7177, -74.0060)
		to := *NewCoordinate(40.6385, -74.0060)
		_ = ride.AddSegment(from, to, time)
		price := ride.Calculate()
		assert.Equal(t, 34345.888941971156, price)
	})

	t.Run("should calculate a ride during sunday day time", func(t *testing.T) {
		ride := NewRide()
		time, _ := time.Parse(layout, "2021-03-07T10:00:00")
		from := *NewCoordinate(40.7177, -74.0060)
		to := *NewCoordinate(40.6385, -74.0060)
		_ = ride.AddSegment(from, to, time)
		price := ride.Calculate()
		assert.Equal(t, 25539.25075172214, price)
	})

	t.Run("should calculate a ride during sunday night time", func(t *testing.T) {
		ride := NewRide()
		time, _ := time.Parse(layout, "2021-03-07T23:00:00")

		from := *NewCoordinate(40.7177, -74.0060)
		to := *NewCoordinate(40.6385, -74.0060)
		_ = ride.AddSegment(from, to, time)
		price := ride.Calculate()
		assert.Equal(t, 44033.190951245066, price)
	})

	t.Run("should calculate a ride during day time with minimum price", func(t *testing.T) {
		ride := NewRide()
		time, _ := time.Parse(layout, "2021-03-01T10:00:00")
		from := *NewCoordinate(40.00000000, -74.0060)
		to := *NewCoordinate(40.00000001, -74.0060)
		_ = ride.AddSegment(from, to, time)
		price := ride.Calculate()
		assert.Equal(t, 10.0, price)

	})

	t.Run("should calculate a ride with multiple segments", func(t *testing.T) {
		ride := NewRide()
		time, _ := time.Parse(layout, "2021-03-01T10:00:00")
		_ = ride.AddSegment(
			*NewCoordinate(40.7177, -74.0060),
			*NewCoordinate(40.6385, -74.0060),
			time,
		)

		_ = ride.AddSegment(
			*NewCoordinate(40.7177, -74.0060),
			*NewCoordinate(40.6385, -74.0060),
			time,
		)
		price := ride.Calculate()
		assert.Equal(t, 36987.88039904586, price)
	})
}
