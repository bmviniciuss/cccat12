package entities

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidCarPlate = errors.New("invalid car plate")
)

type CarPlate struct {
	Value string
}

func isValidCarPlate(value string) bool {
	r, _ := regexp.Compile("^[a-z]{3}[0-9]{4}$")
	return r.MatchString(strings.ToLower(value))
}

func NewCarPlate(value string) (*CarPlate, error) {
	if isValid := isValidCarPlate(value); !isValid {
		return nil, ErrInvalidCarPlate
	}

	return &CarPlate{
		Value: value,
	}, nil
}
