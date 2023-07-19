package entities

import "github.com/google/uuid"

type Passenger struct {
	ID       uuid.UUID
	Name     string
	Email    Email
	Document CPF
}

func CreatePassenger(name, email, document string) (*Passenger, error) {
	cpf, err := NewCPF(document)
	if err != nil {
		return nil, err
	}

	emailVO, err := NewEmail(email)
	if err != nil {
		return nil, err
	}

	p := &Passenger{
		ID:       NewULID(),
		Name:     name,
		Email:    *emailVO,
		Document: *cpf,
	}

	return p, nil
}
