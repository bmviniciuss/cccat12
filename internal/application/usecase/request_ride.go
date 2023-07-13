package usecase

import (
	"context"
	"github.com/bmviniciuss/cccat12/internal/application/repository"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

type RequestRide struct {
	rideRepository repository.Ride
}

func NewRequestRide(rideRepository repository.Ride) *RequestRide {
	return &RequestRide{
		rideRepository: rideRepository,
	}
}

type Coordinate struct {
	Latitude  float64
	Longitude float64
}

type RequestRideInput struct {
	PassengerID string
	From        Coordinate
	To          Coordinate
}

type RequestRideOutput struct {
	ID string
}

func (uc *RequestRide) Execute(ctx context.Context, input RequestRideInput) (*RequestRideOutput, error) {
	from := entities.Coordinate{
		Latitude:  input.From.Latitude,
		Longitude: input.From.Longitude,
	}
	to := entities.Coordinate{
		Latitude:  input.To.Latitude,
		Longitude: input.To.Longitude,
	}
	ride := entities.CreateRide(input.PassengerID, from, to)
	err := uc.rideRepository.Request(ctx, ride)
	if err != nil {
		return nil, err
	}

	return &RequestRideOutput{
		ID: ride.ID,
	}, nil
}
