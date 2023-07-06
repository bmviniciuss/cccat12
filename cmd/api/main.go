package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/bmviniciuss/cccat12/internal/adapters/db/connections"
	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest"
	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/handlers"
	"github.com/bmviniciuss/cccat12/internal/adapters/repositories/pg"
	"github.com/bmviniciuss/cccat12/internal/application/usecase"
)

func main() {
	ctx := context.Background()

	pgm := connections.NewPostgresManager()
	err := pgm.Connect(ctx, connections.PostgresConfig{
		Host:     "localhost",
		Port:     "5432",
		User:     "cccar_user",
		Password: "1234",
		Database: "cccar",
	})
	if err != nil {
		panic(err)
	}
	defer pgm.CloseConnection()

	db := pgm.GetConnection()
	rideCalculatorHandler := handlers.NewRideCalculatorHandler()

	passengerRepository := pg.NewPassengerRepository(db)
	createPassengerUseCase := usecase.NewCreatePassenger(passengerRepository)
	getPassengerUseCase := usecase.NewGetPassenger(passengerRepository)
	passengersHandler := handlers.NewPassengerHandler(createPassengerUseCase, getPassengerUseCase)

	driverRepository := pg.NewDriverRepository(db)
	createDriverUseCase := usecase.NewCreateDriver(driverRepository)
	getDriverUseCase := usecase.NewGetDriver(driverRepository)
	driverHandler := handlers.NewDriverHandler(createDriverUseCase, getDriverUseCase)

	cs := rest.NewServer(
		driverHandler,
		passengersHandler,
		rideCalculatorHandler,
	)
	fmt.Println("Server running on port 3000")
	err = http.ListenAndServe(":3000", cs.Build())
	if err != nil {
		panic(err)
	}
}
