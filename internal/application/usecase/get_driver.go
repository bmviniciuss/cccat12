package usecase

import (
	"context"
	"errors"
	"fmt"

	"github.com/bmviniciuss/cccat12/internal/application/repository"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

type GetDriver struct {
	driverRepository repository.Driver
}

func NewGetDriver(driverRepository repository.Driver) *GetDriver {
	return &GetDriver{
		driverRepository: driverRepository,
	}
}

var (
	ErrDriverNotFound = errors.New("usecase: driver not found")
)

func (g *GetDriver) Execute(ctx context.Context, id string) (*entities.Driver, error) {
	fmt.Println("GetDriver.Execute", id)
	driver, err := g.driverRepository.Get(ctx, id)
	if errors.Is(err, repository.ErrorDriverNotFound) {
		return nil, ErrDriverNotFound
	}

	if err != nil {
		fmt.Printf("GetDriver.Execute: %v\n", err)
		return nil, err
	}

	return driver, nil
}
