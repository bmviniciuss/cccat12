package repository

import (
	"context"
	"errors"

	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

var (
	ErrorPassengerNotFound = errors.New("repository: passenger not found")
)

type Passenger interface {
	Create(ctx context.Context, p *entities.Passenger) error
	Get(ctx context.Context, id string) (*entities.Passenger, error)
}
