package mem

import (
	"context"
	"sync"

	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

type PassengerRepository struct {
	passengers map[string]*entities.Passenger
	lock       sync.Mutex
}

func NewPassengerRepository() *PassengerRepository {
	return &PassengerRepository{
		passengers: make(map[string]*entities.Passenger),
	}
}

func (r *PassengerRepository) Create(ctx context.Context, p *entities.Passenger) (err error) {
	r.lock.Lock()
	r.passengers[p.ID.String()] = p
	r.lock.Unlock()
	return nil
}
