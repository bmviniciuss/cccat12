package entities

import (
	"errors"
	"regexp"
	"strings"
)

var (
	ErrInvalidEmail = errors.New("invalid email")
)

type Email struct {
	Value string
}

func isEmailValid(value string) bool {
	regex := `^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$`
	r, _ := regexp.Compile(regex)
	return r.MatchString(strings.ToLower(value))
}

func NewEmail(value string) (*Email, error) {
	if isValid := isEmailValid(value); !isValid {
		return nil, ErrInvalidEmail
	}

	return &Email{
		Value: value,
	}, nil
}
