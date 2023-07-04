package entities

import (
	"errors"
	"fmt"
	"regexp"
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

func NewCPF(value string) (*CPF, error) {
	cleanValue := cleanCPF(value)
	err := validate(cleanValue)
	if err != nil {
		return nil, err
	}
	model := CPF(cleanValue)
	return &model, nil
}

func cleanCPF(value string) string {
	cleanValue := removeSeparators(value)
	cleanValue = extractDigits(cleanValue)
	return cleanValue
}

func removeSeparators(cpf string) string {
	str := strings.ReplaceAll(cpf, ".", "")
	str = strings.ReplaceAll(str, "-", "")
	str = strings.ReplaceAll(str, " ", "")
	return str
}

func extractDigits(value string) string {
	re := regexp.MustCompile(`\d+`)
	parts := re.FindAllString(value, -1)
	return strings.Join(parts, "")
}

func validate(value string) error {
	if !isInvalidLength(value) {
		return ErrCPFInvalidLength
	}
	vd := extractVerificationDigits(value)
	d1 := calculateDigit(value, 10)
	d2 := calculateDigit(value, 11)

	if vd != fmt.Sprintf("%d%d", d1, d2) {
		return ErrCPFInvalid
	}

	return nil
}

func isInvalidLength(cpf string) bool {
	return len(cpf) == 11
}

func calculateDigit(value string, factor int) int {
	sum := 0
	for _, digit := range value {
		if factor < 2 {
			break
		}

		digit, _ := strconv.Atoi(string(digit))
		sum += digit * factor
		factor--
	}

	rest := sum % 11
	if rest < 2 {
		return 0
	}
	return 11 - rest
}

func extractVerificationDigits(value string) string {
	return value[9:]
}
