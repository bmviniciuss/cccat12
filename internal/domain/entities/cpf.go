package entities

import (
	"errors"
	"strings"
)

type CPF string

var (
	ErrCPFInvalid = errors.New("CPF is invalid")
)

func NewCPF(cpf string) (*CPF, error) {
	cleanCPF := cleanDigits(cpf)
	c := CPF(cleanCPF)
	err := c.Validate()
	if err != nil {
		return nil, err
	}
	return &c, nil
}

func (c CPF) String() string {
	return string(c)
}

func (cpf CPF) Validate() error {
	if len(cpf) != 11 {
		return ErrCPFInvalid
	}
	d1 := calculateDigit(cpf.String()[:9], 10)
	if d1 != int(cpf.String()[9]) {
		return ErrCPFInvalid
	}

	d2 := calculateDigit(cpf.String()[:10], 11)
	if d2 != int(cpf.String()[10]) {
		return ErrCPFInvalid
	}

	return nil
}

func cleanDigits(cpf string) string {
	str := strings.ReplaceAll(cpf, ".", "")
	str = strings.ReplaceAll(str, "-", "")
	str = strings.ReplaceAll(str, " ", "")
	return str
}

func calculateDigit(substr string, multiplierStart int) int {
	multiplier := multiplierStart
	sum := 0

	for i := 0; i < len(substr); i++ {
		digit := int(substr[i])
		sum += digit * multiplier
		multiplier--
	}

	rest := sum % 11
	if rest < 2 {
		return 0
	}
	return 11 - rest
}
