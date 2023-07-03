package usecase

import (
	"context"

	"github.com/bmviniciuss/cccat12/internal/application/repository"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

type CreateDriver struct {
	driverRepository repository.Driver
}

func NewCreateDriver(driverRepository repository.Driver) *CreateDriver {
	return &CreateDriver{
		driverRepository: driverRepository,
	}
}

type CreateDriverInput struct {
	Name        string
	Document    string
	PlateNumber string
}

type CreateDriverOutput struct {
	ID string
}

func (c *CreateDriver) Execute(ctx context.Context, input CreateDriverInput) (*CreateDriverOutput, error) {
	driver, err := entities.CreateDriver(input.Name, input.Document, input.PlateNumber)
	if err != nil {
		return nil, err
	}

	err = c.driverRepository.Create(ctx, driver)
	if err != nil {
		return nil, err
	}

	return &CreateDriverOutput{
		ID: driver.ID.String(),
	}, nil
}
