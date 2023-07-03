package rest

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/handlers"
	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/middlewares"
	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/presentation"
	"github.com/bmviniciuss/cccat12/internal/adapters/repositories/mem"
	"github.com/bmviniciuss/cccat12/internal/application/usecase"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
	"github.com/stretchr/testify/assert"
)

// import (
// 	"encoding/json"
// 	"io"
// 	"net/http/httptest"
// 	"strings"
// 	"testing"

// 	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/handlers"
// 	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/presentation"
// 	"github.com/bmviniciuss/cccat12/internal/adapters/repositories/mem"
// 	"github.com/bmviniciuss/cccat12/internal/application/usecase"
// 	"github.com/gofiber/fiber/v2"
// 	"github.com/stretchr/testify/assert"
// )

type mockDriverHandlers struct{}

func (m *mockDriverHandlers) Create(w http.ResponseWriter, r *http.Request) {}

type mockPassagerHandlers struct{}

func (m *mockPassagerHandlers) Create(w http.ResponseWriter, r *http.Request) {}

type mockRideCalculatorHandlers struct{}

func (m *mockRideCalculatorHandlers) Calculate(w http.ResponseWriter, r *http.Request) {}

func buildMapResponse(in map[string]interface{}) string {
	b, _ := json.Marshal(in)
	return cleanString(string(b))
}

func cleanString(s string) string {
	s = strings.ReplaceAll(s, "\t", "")
	s = strings.ReplaceAll(s, "\n", "")
	return strings.TrimSpace(s)
}

func Test_CalculateRide(t *testing.T) {
	t.Run("should return a ride price", func(t *testing.T) {
		mux := NewServer(
			&mockDriverHandlers{},
			&mockPassagerHandlers{},
			handlers.NewRideCalculatorHandler(),
		).Build()

		req := httptest.NewRequest("POST", "/calculate_ride", strings.NewReader(`{
			"segments": [
				{ "distance": 10,	"date": "2021-03-01T10:00:00" }
			]
		}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		resBody := strings.TrimSpace(rec.Body.String())
		assert.Equal(t, rec.Code, 200)
		assert.Equal(t, buildMapResponse(
			map[string]interface{}{
				"price": 21,
			},
		), resBody)
	})

	t.Run("should return a 422 response if date is not valid", func(t *testing.T) {

		mux := NewServer(
			&mockDriverHandlers{},
			&mockPassagerHandlers{},
			handlers.NewRideCalculatorHandler(),
		).Build()

		req := httptest.NewRequest("POST", "/calculate_ride", strings.NewReader(`{
			"segments": [
				{ "distance": 10,	"date": "2021-03-0110:00:00" }
			]
		}`))
		req.Header.Set("Content-Type", "application/json")
		id := entities.NewULID().String()
		req.Header.Set(middlewares.RequestIDHeader, id)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		resBody := strings.TrimSpace(rec.Body.String())
		assert.Equal(t, rec.Code, 422)
		assert.Equal(t,
			buildMapResponse(map[string]interface{}{
				"id":      id,
				"message": "invalid date",
			}),
			resBody,
		)
	})
}

func Test_CreatePassager(t *testing.T) {
	t.Run("should return a response with passager id", func(t *testing.T) {
		memRepo := mem.NewPassagerRepository()
		usecase := usecase.NewCreatePassager(memRepo)
		mux := NewServer(
			&mockDriverHandlers{},
			handlers.NewPassagerHandler(usecase),
			&mockRideCalculatorHandlers{},
		).Build()

		req := httptest.NewRequest("POST", "/passagers", strings.NewReader(`{
			"name": "Vinicius Barbosa de Medeiros",
			"email": "test@test.com",
			"document": "46021430085"
		}`))
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		resBody := rec.Body.Bytes()
		assert.Equal(t, rec.Code, 201)

		var res presentation.CreatePassagerOutput
		err := json.Unmarshal(resBody, &res)
		assert.NoError(t, err)
		assert.NotEmpty(t, res.ID)
	})

	t.Run("should return an error response with invalid document", func(t *testing.T) {
		memRepo := mem.NewPassagerRepository()
		usecase := usecase.NewCreatePassager(memRepo)
		mux := NewServer(
			&mockDriverHandlers{},
			handlers.NewPassagerHandler(usecase),
			&mockRideCalculatorHandlers{},
		).Build()

		req := httptest.NewRequest("POST", "/passagers", strings.NewReader(`{
			"name": "Vinicius Barbosa de Medeiros",
			"email": "test@test.com",
			"document": "46021430086"
		}`))
		req.Header.Set("Content-Type", "application/json")
		id := entities.NewULID().String()
		req.Header.Set(middlewares.RequestIDHeader, id)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		resBody := cleanString(rec.Body.String())
		assert.Equal(t, rec.Code, 422)

		assert.Equal(t,
			buildMapResponse(map[string]interface{}{
				"id":      id,
				"message": "CPF is invalid",
			}),
			resBody,
		)
	})
}
