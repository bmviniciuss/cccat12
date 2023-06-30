package entities

import "github.com/google/uuid"

type Passager struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Document CPF
}

func CreatePassager(name, email, document string) (*Passager, error) {
	cpf, err := NewCPF(document)
	if err != nil {
		return nil, err
	}

	p := &Passager{
		ID:       NewULID(),
		Name:     name,
		Email:    email,
		Document: *cpf,
	}

	return p, nil
}
