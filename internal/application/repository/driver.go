package repository

import (
	"context"

	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

type Driver interface {
	Create(ctx context.Context, p *entities.Driver) error
}
