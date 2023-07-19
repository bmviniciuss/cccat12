package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmail_New(t *testing.T) {
	t.Run("should return an error for a invalid email", func(t *testing.T) {
		email, err := NewEmail("john_doe@gmail")
		assert.Nil(t, email)
		assert.Equal(t, ErrInvalidEmail, err)
	})

	t.Run("should return email", func(t *testing.T) {
		email, err := NewEmail("john_doe@gmail.com")
		assert.Nil(t, err)
		assert.NotEmpty(t, email.Value)
		assert.Equal(t, email.Value, "john_doe@gmail.com")
	})

}
