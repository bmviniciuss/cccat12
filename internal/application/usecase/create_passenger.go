package usecase

import (
	"context"

	"github.com/bmviniciuss/cccat12/internal/application/repository"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

type CreatePassenger struct {
	passengerRepository repository.Passenger
}

func NewCreatePassenger(passengerRepository repository.Passenger) *CreatePassenger {
	return &CreatePassenger{
		passengerRepository: passengerRepository,
	}
}

func (c *CreatePassenger) Execute(ctx context.Context, input CreatePassengerInput) (*CreatePassengerOutput, error) {
	passenger, err := entities.CreatePassenger(input.Name, input.Email, input.Document)
	if err != nil {
		return nil, err
	}

	err = c.passengerRepository.Create(ctx, passenger)
	if err != nil {
		return nil, err
	}

	return &CreatePassengerOutput{
		ID: passenger.ID.String(),
	}, nil
}

type CreatePassengerInput struct {
	Name     string
	Email    string
	Document string
}

type CreatePassengerOutput struct {
	ID string
}
