package rest

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/bmviniciuss/cccat12/internal/adapters/db/connections"
	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/handlers"
	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/middlewares"
	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/presentation"
	"github.com/bmviniciuss/cccat12/internal/adapters/repositories/mem"
	"github.com/bmviniciuss/cccat12/internal/adapters/repositories/pg"
	"github.com/bmviniciuss/cccat12/internal/application/usecase"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
	"github.com/bmviniciuss/cccat12/testutils"
	"github.com/stretchr/testify/assert"
)

type mockDriverHandlers struct{}

func (m *mockDriverHandlers) Create(w http.ResponseWriter, r *http.Request) {}

type mockPassengerHandlers struct{}

func (m *mockPassengerHandlers) Create(w http.ResponseWriter, r *http.Request) {}

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

func Test_NotFound(t *testing.T) {
	t.Run("should return 404 response", func(t *testing.T) {
		mux := NewServer(
			&mockDriverHandlers{},
			&mockPassengerHandlers{},
			&mockRideCalculatorHandlers{},
		).Build()
		req := httptest.NewRequest(http.MethodPost, "/not_found", nil)
		id := entities.NewULID().String()
		req.Header.Set(middlewares.RequestIDHeader, id)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		resBody := strings.TrimSpace(rec.Body.String())
		assert.Equal(t, rec.Code, 404)
		assert.Equal(t, buildMapResponse(
			map[string]interface{}{
				"id":      id,
				"message": "Not Found",
			},
		), resBody)
	})
}

func Test_CalculateRide(t *testing.T) {
	t.Run("should return a ride price", func(t *testing.T) {
		mux := NewServer(
			&mockDriverHandlers{},
			&mockPassengerHandlers{},
			handlers.NewRideCalculatorHandler(),
		).Build()

		req := httptest.NewRequest(
			http.MethodPost,
			"/calculate_ride",
			strings.NewReader(
				testutils.StringFromMap(
					map[string]interface{}{
						"segments": []map[string]interface{}{
							{
								"distance": 10,
								"date":     "2021-03-01T10:00:00",
							},
						},
					},
				),
			),
		)
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
			&mockPassengerHandlers{},
			handlers.NewRideCalculatorHandler(),
		).Build()

		req := httptest.NewRequest(
			http.MethodPost,
			"/calculate_ride",
			strings.NewReader(
				testutils.StringFromMap(
					map[string]interface{}{
						"segments": []map[string]interface{}{
							{
								"distance": 10,
								"date":     "2021-03-0110:00:00",
							},
						},
					},
				),
			),
		)

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

func Test_CreatePassenger(t *testing.T) {
	t.Run("should return a response with passenger id", func(t *testing.T) {
		memRepo := mem.NewPassengerRepository()
		usecase := usecase.NewCreatePassenger(memRepo)
		mux := NewServer(
			&mockDriverHandlers{},
			handlers.NewPassengerHandler(usecase),
			&mockRideCalculatorHandlers{},
		).Build()

		req := httptest.NewRequest(
			http.MethodPost,
			"/passengers",
			strings.NewReader(
				testutils.StringFromMap(
					map[string]interface{}{
						"name":     "Vinicius Barbosa de Medeiros",
						"email":    "test@test.com",
						"document": testutils.GenerateRandomCPF(),
					},
				),
			),
		)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		resBody := rec.Body.Bytes()
		assert.Equal(t, rec.Code, 201)

		var res presentation.CreatePassengerOutput
		err := json.Unmarshal(resBody, &res)
		assert.NoError(t, err)
		assert.NotEmpty(t, res.ID)
	})

	t.Run("should return an error response with invalid document", func(t *testing.T) {
		memRepo := mem.NewPassengerRepository()
		usecase := usecase.NewCreatePassenger(memRepo)
		mux := NewServer(
			&mockDriverHandlers{},
			handlers.NewPassengerHandler(usecase),
			&mockRideCalculatorHandlers{},
		).Build()

		req := httptest.NewRequest(
			http.MethodPost,
			"/passengers",
			strings.NewReader(
				testutils.StringFromMap(
					map[string]interface{}{
						"name":     "Vinicius Barbosa de Medeiros",
						"email":    "test@test.com",
						"document": "11111111111",
					},
				),
			),
		)
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

func Test_CreateDriver(t *testing.T) {
	ctx := context.Background()
	pgm := connections.NewPostgresManager()

	err := pgm.Connect(ctx, connections.PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "cccar_user",
		Password: "1234",
		Database: "cccar",
	})

	db := pgm.GetConnection()

	if err != nil {
		t.Error(err)
		return
	}

	_, err = db.Exec("DELETE FROM cccar.drivers where true")
	if err != nil {
		t.Error(err)
		return
	}

	t.Cleanup(func() {
		pgm.CloseConnection()
	})
	t.Run("should return a response with driver id", func(t *testing.T) {
		driverRepo := pg.NewDriverRepository(db)
		usecase := usecase.NewCreateDriver(driverRepo)
		mux := NewServer(
			handlers.NewDriverHandler(usecase),
			&mockPassengerHandlers{},
			&mockRideCalculatorHandlers{},
		).Build()

		req := httptest.NewRequest(
			http.MethodPost,
			"/drivers",
			strings.NewReader(
				testutils.StringFromMap(map[string]interface{}{
					"name":      "Vinicius Barbosa de Medeiros",
					"email":     "test@test.com",
					"document":  testutils.GenerateRandomCPF(),
					"car_plate": "ABC-1234",
				}),
			),
		)
		req.Header.Set("Content-Type", "application/json")
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		assert.Equal(t, 201, rec.Code)
		resBody := rec.Body.Bytes()
		var res presentation.CreateDriverOutput
		err := json.Unmarshal(resBody, &res)
		assert.NoError(t, err)
		assert.NotEmpty(t, res.ID)
	})
}
