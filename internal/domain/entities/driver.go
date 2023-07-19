package entities

import "github.com/google/uuid"

type Driver struct {
	ID       uuid.UUID
	Name     string
	Document CPF
	CarPlate CarPlate // TODO: rules for this ??
}

func CreateDriver(name, document, plateNumber string) (*Driver, error) {
	cpf, err := NewCPF(document)
	if err != nil {
		return nil, err
	}

	carPlate, err := NewCarPlate(plateNumber)
	if err != nil {
		return nil, err
	}

	return &Driver{
		ID:       NewULID(),
		Name:     name,
		CarPlate: *carPlate,
		Document: *cpf,
	}, nil
}
