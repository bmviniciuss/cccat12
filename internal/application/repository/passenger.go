package repository

import (
	"context"

	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

type Passenger interface {
	Create(ctx context.Context, p *entities.Passenger) error
}
