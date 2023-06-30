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

var insertPassagerQuery = `
INSERT INTO cccar.passagers
	(id, "name", email, "document", created_at, updated_at)
VALUES($1, $2, $3, $4, now(), now());
`

func (r *PassagerRepository) Create(ctx context.Context, p *entities.Passager) (err error) {
	stmt, err := r.db.PrepareContext(ctx, insertPassagerQuery)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, p.ID, p.Name, p.Email, p.Document)
	if err != nil {
		return err // TODO: What to do with constraints errors?
	}

	return nil
}
