package repository

import (
	"context"
	"errors"

	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

var (
	ErrorDriverNotFound = errors.New("repository: driver not found")
)

type Driver interface {
	Create(ctx context.Context, p *entities.Driver) error
	Get(ctx context.Context, id string) (*entities.Driver, error)
}
