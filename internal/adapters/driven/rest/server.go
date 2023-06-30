package rest

import (
	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/handlers"
	"github.com/gofiber/fiber/v2"
)

type Server struct {
	RideCalculatorHandler *handlers.RideCalculatorHandler
	PassagerHandler       *handlers.PassagerHandler
}

func NewServer(
	rideCalculatorHandler *handlers.RideCalculatorHandler,
	passagerHandler *handlers.PassagerHandler,
) *Server {
	return &Server{
		RideCalculatorHandler: rideCalculatorHandler,
		PassagerHandler:       passagerHandler,
	}
}

func (s *Server) Build() *fiber.App {
	app := fiber.New()
	app.Post("/calculate_ride", s.RideCalculatorHandler.Handle)
	app.Post("/passagers", s.PassagerHandler.Create)
	return app
}
