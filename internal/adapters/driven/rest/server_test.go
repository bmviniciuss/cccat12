package rest

import (
	"context"
	"encoding/json"
	"fmt"
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
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

type mockDriverHandlers struct{}

func (m *mockDriverHandlers) Create(w http.ResponseWriter, r *http.Request) {}
func (m *mockDriverHandlers) Get(w http.ResponseWriter, r *http.Request)    {}

type mockPassengerHandlers struct{}

func (m *mockPassengerHandlers) Create(w http.ResponseWriter, r *http.Request) {}
func (m *mockPassengerHandlers) Get(w http.ResponseWriter, r *http.Request)    {}

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

func buildTestServer(db *sqlx.DB) *chi.Mux {
	driverRepo := pg.NewDriverRepository(db)
	createDriver := usecase.NewCreateDriver(driverRepo)
	getDriver := usecase.NewGetDriver(driverRepo)

	passengerRepo := pg.NewPassengerRepository(db)
	createPassenger := usecase.NewCreatePassenger(passengerRepo)
	getPassenger := usecase.NewGetPassenger(passengerRepo)

	return NewServer(
		handlers.NewDriverHandler(createDriver, getDriver),
		handlers.NewPassengerHandler(createPassenger, getPassenger),
		handlers.NewRideCalculatorHandler(),
	).Build()
}

func TestApi_CalculateRide(t *testing.T) {
	pgm := withDatabase(t, context.Background())

	t.Cleanup(func() {
		pgm.CloseConnection()
	})

	// ride.AddPosition(-27.584905257808835, -48.545022195325124, time)
	// ride.AddPosition(-27.496887588317275, -48.522234807851476, time)

	t.Run("should return a ride price", func(t *testing.T) {

		mux := buildTestServer(pgm.GetConnection())
		req := httptest.NewRequest(
			http.MethodPost,
			"/calculate_ride",
			strings.NewReader(
				testutils.StringFromMap(
					map[string]interface{}{
						"positions": []map[string]interface{}{
							{
								"lat":  -27.584905257808835,
								"long": -48.545022195325124,
								"date": "2021-03-01T10:00:00",
							},
							{
								"lat":  -27.496887588317275,
								"long": -48.522234807851476,
								"date": "2021-03-01T10:00:00",
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
				"price": 21.08753314776229,
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
						"positions": []map[string]interface{}{
							{
								"from": map[string]float64{
									"lat":  38.898556,
									"long": -77.037852,
								},
								"to": map[string]float64{
									"lat":  38.897147,
									"long": -77.043934,
								},
								"date": "2021-03-0110:00:00",
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
		createPassengerUseCase := usecase.NewCreatePassenger(memRepo)
		getPassenger := usecase.NewGetPassenger(memRepo)
		mux := NewServer(
			&mockDriverHandlers{},
			handlers.NewPassengerHandler(createPassengerUseCase, getPassenger),
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
		createPassengerUseCase := usecase.NewCreatePassenger(memRepo)
		getPassenger := usecase.NewGetPassenger(memRepo)
		mux := NewServer(
			&mockDriverHandlers{},
			handlers.NewPassengerHandler(createPassengerUseCase, getPassenger),
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

func withDatabase(t *testing.T, ctx context.Context) *connections.PostgresManager {
	pgm := connections.NewPostgresManager()

	err := pgm.Connect(ctx, connections.PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "cccar_user",
		Password: "1234",
		Database: "cccar",
	})

	if err != nil {
		t.Error(err)
		return nil
	}

	return pgm
}

func Test_CreateDriver(t *testing.T) {
	pgm := withDatabase(t, context.Background())
	db := pgm.GetConnection()

	_, err := db.Exec("DELETE FROM cccar.drivers where true")
	if err != nil {
		t.Error(err)
		return
	}

	t.Cleanup(func() {
		pgm.CloseConnection()
	})

	t.Run("should return a response with driver id", func(t *testing.T) {
		driverRepo := pg.NewDriverRepository(db)
		createDriver := usecase.NewCreateDriver(driverRepo)
		getDriver := usecase.NewGetDriver(driverRepo)
		mux := NewServer(
			handlers.NewDriverHandler(createDriver, getDriver),
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
					"car_plate": "ABC1234",
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

type driverRow struct {
	ID          string `db:"id"`
	Name        string `db:"name"`
	Document    string `db:"document"`
	PlateNumber string `db:"plate_number"`
}

func Test_GetDriver(t *testing.T) {
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

	t.Cleanup(func() {
		pgm.CloseConnection()
	})
	t.Run("should return a response with driver", func(t *testing.T) {
		driverRepo := pg.NewDriverRepository(db)
		createDriver := usecase.NewCreateDriver(driverRepo)
		getDriver := usecase.NewGetDriver(driverRepo)
		mux := NewServer(
			handlers.NewDriverHandler(createDriver, getDriver),
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
					"car_plate": "ABC1234",
				}),
			),
		)
		rec := httptest.NewRecorder()
		mux.ServeHTTP(rec, req)
		resBody := rec.Body.Bytes()
		var res presentation.CreateDriverOutput
		_ = json.Unmarshal(resBody, &res)

		req = httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf("/drivers/%s", res.ID),
			nil,
		)
		req.Header.Set("Content-Type", "application/json")
		rec = httptest.NewRecorder()
		mux.ServeHTTP(rec, req)

		assert.Equal(t, 200, rec.Code)
	})

	t.Run("should return 404 if driver is not found", func(t *testing.T) {
		driverRepo := pg.NewDriverRepository(db)
		createDriver := usecase.NewCreateDriver(driverRepo)
		getDriver := usecase.NewGetDriver(driverRepo)
		mux := NewServer(
			handlers.NewDriverHandler(createDriver, getDriver),
			&mockPassengerHandlers{},
			&mockRideCalculatorHandlers{},
		).Build()

		rec := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf("/drivers/%s", entities.NewULID().String()),
			nil,
		)
		reqID := entities.NewULID().String()
		req.Header.Set(middlewares.RequestIDHeader, reqID)
		req.Header.Set("Content-Type", "application/json")
		mux.ServeHTTP(rec, req)

		resBody := cleanString(rec.Body.String())

		assert.Equal(t, 404, rec.Code)
		assert.Equal(t,
			buildMapResponse(map[string]interface{}{
				"id":      reqID,
				"message": "Not Found",
			}),
			resBody,
		)
	})
}

func Test_GetPassenger(t *testing.T) {
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

	_, err = db.Exec("DELETE FROM cccar.passengers where true")
	if err != nil {
		t.Error(err)
		return
	}

	if err != nil {
		t.Error(err)
		return
	}

	t.Cleanup(func() {
		pgm.CloseConnection()
	})

	t.Run("should return a response with passenger", func(t *testing.T) {
		mux := buildTestServer(db)
		// create a passenger
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

		var cpo presentation.CreatePassengerOutput
		err := json.Unmarshal(resBody, &cpo)
		if err != nil {
			t.Error(err)
			return
		}

		req2 := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf("/passengers/%s", cpo.ID),
			nil,
		)
		req.Header.Set("Content-Type", "application/json")
		rec2 := httptest.NewRecorder()
		mux.ServeHTTP(rec2, req2)
		assert.Equal(t, 200, rec2.Result().StatusCode)
	})

	t.Run("should return 404 if passenger does not exists", func(t *testing.T) {
		mux := buildTestServer(db)
		randomID := entities.NewULID().String()
		rec := httptest.NewRecorder()
		req := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf("/passengers/%s", randomID),
			nil,
		)
		reqID := entities.NewULID().String()
		req.Header.Set(middlewares.RequestIDHeader, reqID)
		req.Header.Set("Content-Type", "application/json")
		mux.ServeHTTP(rec, req)

		resBody := cleanString(rec.Body.String())
		assert.Equal(t, 404, rec.Code)
		assert.Equal(t,
			buildMapResponse(map[string]interface{}{
				"id":      reqID,
				"message": "Not Found",
			}),
			resBody,
		)
	})
}
