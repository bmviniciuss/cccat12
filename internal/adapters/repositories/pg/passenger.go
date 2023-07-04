package pg

import (
	"context"

	"github.com/bmviniciuss/cccat12/internal/application/repository"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
	"github.com/jmoiron/sqlx"
)

type PassengerRepository struct {
	db *sqlx.DB
}

func NewPassengerRepository(db *sqlx.DB) *PassengerRepository {
	return &PassengerRepository{
		db: db,
	}
}

var (
	_ repository.Passenger = (*PassengerRepository)(nil)
)

var insertPassengerQuery = `
INSERT INTO cccar.passengers
	(id, "name", email, "document", created_at, updated_at)
VALUES($1, $2, $3, $4, now(), now());
`

func (r *PassengerRepository) Create(ctx context.Context, p *entities.Passenger) (err error) {
	stmt, err := r.db.PrepareContext(ctx, insertPassengerQuery)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, p.ID, p.Name, p.Email, p.Document)
	if err != nil {
		return err // TODO: What to do with constraints errors?
	}

	return nil
}
