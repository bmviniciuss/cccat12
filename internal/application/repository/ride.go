package repository

import (
	"context"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

type Ride interface {
	Request(ctx context.Context, r *entities.Ride) error
}
