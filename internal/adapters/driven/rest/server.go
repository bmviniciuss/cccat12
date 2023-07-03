package rest

import (
	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/handlers"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	RideCalculatorHandler *handlers.RideCalculatorHandler
	PassagerHandler       handlers.PassagerHandlerPort
	driverHandlers        handlers.DriverHandlerPort
}

func NewServer(
	rideCalculatorHandler *handlers.RideCalculatorHandler,
	passagerHandler handlers.PassagerHandlerPort,
	driverHandlers handlers.DriverHandlerPort,
) *Server {
	return &Server{
		RideCalculatorHandler: rideCalculatorHandler,
		PassagerHandler:       passagerHandler,
		driverHandlers:        driverHandlers,
	}
}

func (s *Server) Build() *fiber.App {
	app := fiber.New()
	app.Post("/calculate_ride", s.RideCalculatorHandler.Handle)
	app.Post("/passagers", s.PassagerHandler.Create)
	app.Post("/drivers", s.driverHandlers.Create)
	return app
}
