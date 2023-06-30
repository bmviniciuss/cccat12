package repository

import (
	"context"

	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

type Passager interface {
	Create(ctx context.Context, p *entities.Passager) error
}
