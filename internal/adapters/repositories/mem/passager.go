package mem

import (
	"context"
	"sync"

	"github.com/bmviniciuss/cccat12/internal/domain/entities"
)

type PassagerRepository struct {
	passagers map[string]*entities.Passager
	lock      sync.Mutex
}

func NewPassagerRepository() *PassagerRepository {
	return &PassagerRepository{
		passagers: make(map[string]*entities.Passager),
	}
}

func (r *PassagerRepository) Create(ctx context.Context, p *entities.Passager) (err error) {
	r.lock.Lock()
	r.passagers[p.ID.String()] = p
	r.lock.Unlock()
	return nil
}
