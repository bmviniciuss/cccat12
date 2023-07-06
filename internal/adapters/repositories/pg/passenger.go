package pg

import (
	"context"
	"database/sql"
	"errors"

	"github.com/bmviniciuss/cccat12/internal/application/repository"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
	"github.com/google/uuid"
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

type passengerRow struct {
	ID       string `db:"id"`
	Name     string `db:"name"`
	Email    string `db:"email"`
	Document string `db:"document"`
}

var getPassengerQuery = `
SELECT
	id, name, email, document
FROM cccar.passengers
WHERE id = $1;
`

func (r *PassengerRepository) Get(ctx context.Context, id string) (*entities.Passenger, error) {
	stmt, err := r.db.PreparexContext(ctx, getPassengerQuery)
	if err != nil {
		return nil, err
	}

	row := passengerRow{}

	err = stmt.GetContext(ctx, &row, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrorPassengerNotFound
	} else if err != nil {
		return nil, err
	}

	uid, err := uuid.Parse(row.ID)
	if err != nil {
		return nil, err
	}

	document, err := entities.NewCPF(row.Document)
	if err != nil {
		return nil, err
	}

	return &entities.Passenger{
		ID:       uid,
		Name:     row.Name,
		Email:    row.Email,
		Document: *document,
	}, nil
}
