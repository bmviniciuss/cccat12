package entities

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRide_Calculate(t *testing.T) {
	t.Run("should calculate the price of a day ride", func(t *testing.T) {
		time, _ := time.Parse(TimeLayout, "2021-03-01T10:00:00")
		ride := NewRide()
		ride.AddPosition(-27.584905257808835, -48.545022195325124, time)
		ride.AddPosition(-27.496887588317275, -48.522234807851476, time)
		price, err := ride.Calculate()
		assert.Nil(t, err)
		assert.Equal(t, 21.08753314776229, price)
	})

	t.Run("should calculate the price of a night ride", func(t *testing.T) {
		time, _ := time.Parse(TimeLayout, "2021-03-01T23:00:00")
		ride := NewRide()
		ride.AddPosition(-27.584905257808835, -48.545022195325124, time)
		ride.AddPosition(-27.496887588317275, -48.522234807851476, time)
		price, err := ride.Calculate()
		assert.Nil(t, err)
		assert.Equal(t, 39.162561560129966, price)
	})

	t.Run("should calculate the price of a sundays`s day ride", func(t *testing.T) {
		time, _ := time.Parse(TimeLayout, "2021-03-07T10:00:00")
		ride := NewRide()
		ride.AddPosition(-27.584905257808835, -48.545022195325124, time)
		ride.AddPosition(-27.496887588317275, -48.522234807851476, time)
		price, err := ride.Calculate()
		assert.Nil(t, err)
		assert.Equal(t, 29.120879108814588, price)
	})

	t.Run("should calculate the price of a sunday`s day ride", func(t *testing.T) {
		time, _ := time.Parse(TimeLayout, "2021-03-07T23:00:00")
		ride := NewRide()
		ride.AddPosition(-27.584905257808835, -48.545022195325124, time)
		ride.AddPosition(-27.496887588317275, -48.522234807851476, time)
		price, err := ride.Calculate()
		assert.Nil(t, err)
		assert.Equal(t, 50.20841225657688, price)
	})

	t.Run("should return a min price ride", func(t *testing.T) {
		time, _ := time.Parse(TimeLayout, "2021-03-07T23:00:00")
		ride := NewRide()
		ride.AddPosition(-27.584905257808835, -48.545022195325124, time)
		ride.AddPosition(-27.584905257808836, -48.545032195325124, time)
		price, err := ride.Calculate()
		assert.Nil(t, err)
		assert.Equal(t, 10.0, price)
	})
}
