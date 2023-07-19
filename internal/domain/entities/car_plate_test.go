package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarPlate_New(t *testing.T) {
	t.Run("should return an error for a invalid car plate", func(t *testing.T) {
		cp, err := NewCarPlate("AA1234")
		assert.Nil(t, cp)
		assert.Equal(t, ErrInvalidCarPlate, err)
	})

	t.Run("should return car plate", func(t *testing.T) {
		cp, err := NewCarPlate("AAA1234")
		assert.Nil(t, err)
		assert.Equal(t, cp.Value, "AAA1234")
	})
}
