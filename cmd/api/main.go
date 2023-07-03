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

	passagerRepository := pg.NewPassagerRepository(db)
	createPassagerUseCase := usecase.NewCreatePassager(passagerRepository)
	passagersHandler := handlers.NewPassagerHandler(createPassagerUseCase)

	driverRepository := pg.NewDriverRepository(db)
	createDriverUseCase := usecase.NewCreateDriver(driverRepository)
	driverHandler := handlers.NewDriverHandler(createDriverUseCase)

	cs := rest.NewServer(
		driverHandler,
		passagersHandler,
		rideCalculatorHandler,
	)
	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", cs.Build())
}
