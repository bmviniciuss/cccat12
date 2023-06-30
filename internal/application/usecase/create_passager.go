package usecase

import (
	"context"

	"github.com/bmviniciuss/cccat12/internal/application/repository"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

type CreatePassager struct {
	passagerRepository repository.Passager
}

func NewCreatePassager(passagerRepository repository.Passager) *CreatePassager {
	return &CreatePassager{
		passagerRepository: passagerRepository,
	}
}

func (c *CreatePassager) Execute(ctx context.Context, input CreatePassagerInput) (*CreatePassagerOutput, error) {
	passager, err := entities.CreatePassager(input.Name, input.Email, input.Document)
	if err != nil {
		return nil, err
	}

	err = c.passagerRepository.Create(ctx, passager)
	if err != nil {
		return nil, err
	}

	return &CreatePassagerOutput{
		ID: passager.ID.String(),
	}, nil
}

type CreatePassagerInput struct {
	Name     string
	Email    string
	Document string
}

type CreatePassagerOutput struct {
	ID string
}
