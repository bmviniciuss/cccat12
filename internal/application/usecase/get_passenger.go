package usecase

import (
	"context"
	"errors"

	"github.com/bmviniciuss/cccat12/internal/application/repository"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

type GetPassenger struct {
	passengerRepository repository.Passenger
}

func NewGetPassenger(passengerRepository repository.Passenger) *GetPassenger {
	return &GetPassenger{
		passengerRepository: passengerRepository,
	}
}

var (
	ErrorPassengerNotFound = errors.New("usecase: passenger not found")
)

func (uc *GetPassenger) Execute(ctx context.Context, id string) (*entities.Passenger, error) {
	passenger, err := uc.passengerRepository.Get(ctx, id)
	if errors.Is(err, repository.ErrorPassengerNotFound) {
		return nil, ErrorPassengerNotFound
	}
	if err != nil {
		return nil, err
	}

	return passenger, nil
}
