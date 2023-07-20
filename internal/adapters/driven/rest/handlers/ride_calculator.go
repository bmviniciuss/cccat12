package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/presentation"
	"github.com/bmviniciuss/cccat12/internal/application/usecase"
	"github.com/bmviniciuss/cccat12/internal/customcontext"
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
	useCase := usecase.NewCalculateRide()
	useCaseInput := usecase.CalculateRideInput{}
	positions := make([]usecase.CalculateRidePosition, len(input.Positions))
	for i, pos := range input.Positions {
		positions[i] = usecase.CalculateRidePosition{
			Lat:  pos.Lat,
			Long: pos.Long,
			Date: pos.Date,
		}
	}
	useCaseInput.Positions = positions

	out, err := useCase.Execute(useCaseInput)
	if err != nil {
		render.Render(w, r, presentation.ErrUnprocessableEntity(reqID, err))
		return
	}

	res := &presentation.CalculateRideOutput{
		Price: out,
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, res)
}
