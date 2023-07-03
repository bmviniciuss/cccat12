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

type DriverHandler struct {
	createDriver *usecase.CreateDriver
}

func NewDriverHandler(createDriver *usecase.CreateDriver) *DriverHandler {
	return &DriverHandler{
		createDriver: createDriver,
	}
}

var (
	_ ports.DriverHandlersPort = (*DriverHandler)(nil)
)

func (h *DriverHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqID, _ := customcontext.RequestID(ctx)
	var input presentation.CreateDriverInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		render.Render(w, r, presentation.ErrBadRequest(reqID, err))
		return
	}
	// TODO: validation
	res, err := h.createDriver.Execute(ctx, usecase.CreateDriverInput{
		Name:        input.Name,
		Document:    input.Document,
		PlateNumber: input.PlateNumber,
	})
	if err != nil {
		render.Render(w, r, presentation.ErrUnprocessableEntity(reqID, err))
	}
	out := &presentation.CreateDriverOutput{
		ID: res.ID,
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, out)
}
