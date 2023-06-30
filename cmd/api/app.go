package main

import (
	"github.com/bmviniciuss/cccat12/cmd/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func Build() *fiber.App {
	app := fiber.New()

	app.Post("/calculate_ride", handlers.CalculateRideHandler)
	passager := app.Group("/passagers")
	passager.Post("/", handlers.CreatePassagerHandler)
	return app
}
