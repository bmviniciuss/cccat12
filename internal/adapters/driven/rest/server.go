package rest

import (
	"net/http"

	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/middlewares"
	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/presentation"
	"github.com/bmviniciuss/cccat12/internal/customcontext"
	"github.com/bmviniciuss/cccat12/internal/ports"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Server struct {
	driverHandlers         ports.DriverHandlersPort
	passagerHandlers       ports.PassagerHandlersPort
	rideCalculatorHandlers ports.RideCalculatorHandlersPort
}

func NewServer(
	driverHandlers ports.DriverHandlersPort,
	passagerHandlers ports.PassagerHandlersPort,
	rideCalculatorHandlers ports.RideCalculatorHandlersPort,
) *Server {
	return &Server{
		driverHandlers:         driverHandlers,
		passagerHandlers:       passagerHandlers,
		rideCalculatorHandlers: rideCalculatorHandlers,
	}
}

func (s *Server) Build() *chi.Mux {
	r := chi.NewRouter()
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middlewares.RequestID)
	r.Use(middleware.Logger)

	r.Post("/passagers", s.passagerHandlers.Create)
	r.Post("/passagers", s.passagerHandlers.Create)
	r.Post("/calculate_ride", s.rideCalculatorHandlers.Calculate)

	r.NotFound(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		reqID, _ := customcontext.RequestID(ctx)
		render.Render(w, r, presentation.ErrNotFound(reqID, nil))
	})

	return r
}
