package pg

import (
	"context"

	"github.com/bmviniciuss/cccat12/internal/application/repository"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
	"github.com/jmoiron/sqlx"
)

type DriverRepository struct {
	db *sqlx.DB
}

func NewDriverRepository(db *sqlx.DB) *DriverRepository {
	return &DriverRepository{
		db: db,
	}
}

var (
	_ repository.Driver = (*DriverRepository)(nil)
)

var insertDriverQuery = `
INSERT INTO cccar.drivers
	(id, "name", "document", plate_number, created_at, updated_at)
	VALUES
	($1, $2, $3, $4, now(), now());
`

func (r *DriverRepository) Create(ctx context.Context, driver *entities.Driver) (err error) {
	stmt, err := r.db.PrepareContext(ctx, insertDriverQuery)
	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, driver.ID, driver.Name, driver.Document, driver.PlateNumber)
	if err != nil {
		return err // TODO: What to do with constraints errors?
	}

	return nil
}
