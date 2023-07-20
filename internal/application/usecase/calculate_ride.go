package usecase

import (
	"time"

	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

type CalculateRide struct {
}

func NewCalculateRide() *CalculateRide {
	return &CalculateRide{}
}

type CalculateRidePosition struct {
	Lat  float64
	Long float64
	Date string
}

type CalculateRideInput struct {
	Positions []CalculateRidePosition
}

type CalculateRideOutput struct {
	Price float64
}

func (c CalculateRide) Execute(input CalculateRideInput) (float64, error) {
	ride := entities.NewRide()

	for _, pos := range input.Positions {
		date, err := time.Parse(entities.TimeLayout, pos.Date)
		if err != nil {
			return 0, entities.ErrInvalidDate
		}
		ride.AddPosition(pos.Lat, pos.Long, date)
	}

	price, err := ride.Calculate()
	if err != nil {
		return 0, err
	}

	return price, nil
}
