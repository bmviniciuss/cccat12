package pg

import (
	"context"

	"github.com/bmviniciuss/cccat12/internal/application/repository"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
	"github.com/jmoiron/sqlx"
)

type PassagerRepository struct {
	db *sqlx.DB
}

func NewPassagerRepository(db *sqlx.DB) *PassagerRepository {
	return &PassagerRepository{
		db: db,
	}
}

var (
	_ repository.Passager = (*PassagerRepository)(nil)
)

func (r *PassagerRepository) Create(ctx context.Context, p *entities.Passager) error {
	return nil
}
