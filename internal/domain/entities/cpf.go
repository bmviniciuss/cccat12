package entities

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type CPF string

func (c CPF) String() string {
	return string(c)
}

var (
	ErrCPFInvalid       = errors.New("CPF is invalid")
	ErrCPFInvalidLength = fmt.Errorf("CPF must have 11 characters")
	ErrCPFNonDigit      = errors.New("CPF contains non digit characters")
)

func NewCPF(cpf string) (*CPF, error) {
	cleanCPF := removeSeparators(cpf)
	if !isNumericString(cleanCPF) {
		return nil, ErrCPFNonDigit
	}
	model := CPF(cleanCPF)
	err := model.Validate()
	if err != nil {
		return nil, err
	}
	return &model, nil
}

func removeSeparators(cpf string) string {
	str := strings.ReplaceAll(cpf, ".", "")
	str = strings.ReplaceAll(str, "-", "")
	str = strings.ReplaceAll(str, " ", "")
	return str
}

func isNumericString(str string) bool {
	_, err := strconv.Atoi(str)
	return err == nil
}

func (cpf CPF) Validate() error {
	if len(cpf) != 11 {
		return ErrCPFInvalidLength
	}

	err := cpf.validateDigit(9, string(cpf.String()[9]), 10)
	if err != nil {
		return ErrCPFInvalid
	}

	err = cpf.validateDigit(10, string(cpf.String()[10]), 11)
	if err != nil {
		return ErrCPFInvalid
	}

	return nil
}

func (cpf CPF) validateDigit(offset int, digit string, multiplier int) error {
	computedD, err := calculateDigit(cpf.String()[:offset], multiplier)
	if err != nil {
		return err
	}

	d, err := strconv.Atoi(digit)
	if err != nil {
		return err
	}

	if computedD != d {
		return ErrCPFInvalid
	}

	return nil
}

func calculateDigit(substr string, multiplierStart int) (int, error) {
	sum := 0
	fmt.Println()
	for i := 0; i < len(substr); i++ {
		digit, err := strconv.Atoi(string(substr[i]))
		if err != nil {
			return -1, nil
		}
		sum += digit * (multiplierStart - i)
		fmt.Printf("%d * %d = %d\n", digit, multiplierStart-i, digit*(multiplierStart-i))
	}
	fmt.Printf("sum: %d\n", sum)
	rest := sum % 11
	fmt.Printf("rest: %d\n", rest)
	if rest < 2 {
		return 0, nil
	}
	return 11 - rest, nil
}
