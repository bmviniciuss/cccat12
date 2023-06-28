package main

import (
	"github.com/bmviniciuss/cccat12/cmd/api/handlers"
	"github.com/gofiber/fiber/v2"
)

func Build() *fiber.App {
	app := fiber.New()

	app.Post("/calculate_ride", handlers.CalculateRideHandler)
	return app
}
