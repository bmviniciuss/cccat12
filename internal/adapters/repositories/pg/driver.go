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

	_, err = stmt.ExecContext(ctx,
		driver.ID, driver.Name, driver.Document, driver.CarPlate.Value,
	)
	if err != nil {
		return err // TODO: What to do with constraints errors?
	}

	return nil
}

var getDriverQuery = `
SELECT 
	id, "name", "document", plate_number
FROM cccar.drivers
WHERE id = $1;
`

type driverRow struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Document    string `db:"document"`
	PlateNumber string `db:"plate_number"`
}

func (r *DriverRepository) Get(ctx context.Context, id string) (*entities.Driver, error) {
	stmt, err := r.db.PreparexContext(ctx, getDriverQuery)
	if err != nil {
		return nil, err
	}

	var row driverRow
	err = stmt.QueryRowxContext(ctx, id).StructScan(&row)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, repository.ErrorDriverNotFound
	}
	if err != nil {
		return nil, err
	}

	uuid, err := uuid.Parse(row.ID)
	if err != nil {
		return nil, err
	}

	document, err := entities.NewCPF(row.Document)
	if err != nil {
		return nil, err
	}

	cp, err := entities.NewCarPlate(row.PlateNumber)
	if err != nil {
		return nil, err
	}

	return &entities.Driver{
		ID:       uuid,
		Name:     row.Name,
		Document: *document,
		CarPlate: *cp,
	}, nil
}
