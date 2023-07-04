package entities

import "github.com/google/uuid"

type Passenger struct {
	ID       uuid.UUID
	Name     string
	Email    string
	Document CPF
}

func CreatePassenger(name, email, document string) (*Passenger, error) {
	cpf, err := NewCPF(document)
	if err != nil {
		return nil, err
	}

	p := &Passenger{
		ID:       NewULID(),
		Name:     name,
		Email:    email,
		Document: *cpf,
	}

	return p, nil
}
