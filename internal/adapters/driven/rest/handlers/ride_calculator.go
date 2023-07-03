package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/presentation"
	"github.com/bmviniciuss/cccat12/internal/customcontext"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
	"github.com/bmviniciuss/cccat12/internal/ports"
	"github.com/go-chi/render"
)

type RideCalculatorHandler struct{}

func NewRideCalculatorHandler() *RideCalculatorHandler {
	return &RideCalculatorHandler{}
}

var (
	_ ports.RideCalculatorHandlersPort = (*RideCalculatorHandler)(nil)
)

func (h *RideCalculatorHandler) Calculate(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqID, _ := customcontext.RequestID(ctx)

	var input presentation.CalculateRideInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		render.Render(w, r, presentation.ErrBadRequest(reqID, err))
		return
	}

	// TODO: validation

	ride := entities.NewRide()
	for _, segment := range input.Segments {
		time, err := time.Parse(entities.TimeLayout, segment.Date)
		if err != nil {
			render.Render(w, r, presentation.ErrUnprocessableEntity(reqID, errors.New("invalid date")))
			return
		}
		err = ride.AddSegment(segment.Distance, time)
		if err != nil {
			render.Render(w, r, presentation.ErrUnprocessableEntity(reqID, err))
			return
		}
	}

	price := ride.Calculate()
	res := &presentation.CalculateRideOutput{
		Price: price,
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, res)
}
