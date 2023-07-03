package rest

import (
	"encoding/json"
	"io"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/handlers"
	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/presentation"
	"github.com/bmviniciuss/cccat12/internal/adapters/repositories/mem"
	"github.com/bmviniciuss/cccat12/internal/application/usecase"
	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

type handlerMock struct{}

func (m *handlerMock) Create(c *fiber.Ctx) error {
	return nil
}

func Test_CalculateRide(t *testing.T) {
	t.Run("should return a ride price", func(t *testing.T) {
		app := NewServer(
			handlers.NewRideCalculatorHandler(),
			&handlerMock{},
			&handlerMock{},
		).Build()

		req := httptest.NewRequest("POST", "/calculate_ride", strings.NewReader(`{
			"segments": [
				{ "distance": 10,	"date": "2021-03-01T10:00:00" }
			]
		}`))
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req, -1)
		assert.NoError(t, err)
		resBody, _ := io.ReadAll(res.Body)
		defer res.Body.Close()
		assert.Equal(t, res.StatusCode, 200)
		assert.Equal(t, `{"price":21}`, string(resBody))
	})

	t.Run("should return a 422 response if date is not valid", func(t *testing.T) {
		app := NewServer(
			handlers.NewRideCalculatorHandler(),
			&handlerMock{},
			&handlerMock{},
		).Build()
		req := httptest.NewRequest("POST", "/calculate_ride", strings.NewReader(`{
			"segments": [
				{ "distance": 10,	"date": "2021-03-0110:00:00" }
			]
		}`))
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req, -1)
		assert.NoError(t, err)
		resBody, _ := io.ReadAll(res.Body)
		defer res.Body.Close()
		assert.Equal(t, res.StatusCode, 422)
		assert.Equal(t, `{"message":"Invalid Date"}`, string(resBody))
	})
}

func Test_CreatePassager(t *testing.T) {
	t.Run("should return a response with passager id", func(t *testing.T) {
		memRepo := mem.NewPassagerRepository()
		usecase := usecase.NewCreatePassager(memRepo)
		app := NewServer(
			handlers.NewRideCalculatorHandler(),
			handlers.NewPassagerHandler(usecase),
			&handlerMock{},
		).Build()

		req := httptest.NewRequest("POST", "/passagers", strings.NewReader(`{
			"name": "Vinicius Barbosa de Medeiros",
			"email": "test@test.com",
			"document": "46021430085"
		}`))
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req, -1)
		assert.NoError(t, err)
		assert.Equal(t, res.StatusCode, 200)
		resBody, _ := io.ReadAll(res.Body)
		defer res.Body.Close()

		var out presentation.CreatePassagerOutput
		_ = json.Unmarshal(resBody, &out)
		assert.NotNil(t, out)
		assert.NotNil(t, out.ID)
	})

	t.Run("should return an error response with invalid document", func(t *testing.T) {
		memRepo := mem.NewPassagerRepository()
		usecase := usecase.NewCreatePassager(memRepo)
		app := NewServer(
			handlers.NewRideCalculatorHandler(),
			handlers.NewPassagerHandler(usecase),
			&handlerMock{},
		).Build()

		req := httptest.NewRequest("POST", "/passagers", strings.NewReader(`{
			"name": "Vinicius Barbosa de Medeiros",
			"email": "test@test.com",
			"document": "46021430086"
		}`))
		req.Header.Set("Content-Type", "application/json")

		res, err := app.Test(req, -1)
		assert.NoError(t, err)
		resBody, _ := io.ReadAll(res.Body)
		defer res.Body.Close()

		assert.Equal(t, res.StatusCode, 422)
		assert.Equal(t, `{"message":"CPF is invalid"}`, string(resBody))
	})
}
