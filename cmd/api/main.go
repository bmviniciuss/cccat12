package main

import (
	"context"

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

	server := rest.NewServer(
		rideCalculatorHandler,
		passagersHandler,
		driverHandler,
	)
	app := server.Build()
	app.Listen(":3000")
}
