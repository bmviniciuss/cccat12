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
		err := ride.AddSegment(0, time.Now())
		assert.Equal(t, err, ErrInvalidSegmentDistance)
	})

	t.Run("should calcule a ride during day time", func(t *testing.T) {
		r := NewRide()
		time, _ := time.Parse(layout, "2021-03-01T10:00:00")
		r.AddSegment(10, time)
		price := r.Calculate()
		assert.Equal(t, price, 21.0)
	})

	t.Run("should calcule a ride during night time", func(t *testing.T) {
		r := NewRide()
		time, _ := time.Parse(layout, "2021-03-01T23:00:00")
		r.AddSegment(10, time)
		price := r.Calculate()
		assert.Equal(t, price, 39.0)
	})

	t.Run("should calcule a ride during sunday day time", func(t *testing.T) {
		r := NewRide()
		time, _ := time.Parse(layout, "2021-03-07T10:00:00")
		r.AddSegment(10, time)
		price := r.Calculate()
		assert.Equal(t, price, 29.0)
	})

	t.Run("should calcule a ride during sunday night time", func(t *testing.T) {
		r := NewRide()
		time, _ := time.Parse(layout, "2021-03-07T23:00:00")
		r.AddSegment(10, time)
		price := r.Calculate()
		assert.Equal(t, price, 50.0)
	})

	t.Run("should calcule a ride during day time with minimum price", func(t *testing.T) {
		r := NewRide()
		time, _ := time.Parse(layout, "2021-03-01T10:00:00")
		r.AddSegment(3, time)
		price := r.Calculate()
		assert.Equal(t, price, 10.0)
	})

	t.Run("should calculate a ride wuth multiple segments", func(t *testing.T) {
		r := NewRide()
		time, _ := time.Parse(layout, "2021-03-01T10:00:00")
		r.AddSegment(10, time)
		r.AddSegment(10, time)
		price := r.Calculate()
		assert.Equal(t, price, 42.0)
	})
}
