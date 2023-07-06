package handlers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/presentation"
	"github.com/bmviniciuss/cccat12/internal/application/usecase"
	"github.com/bmviniciuss/cccat12/internal/customcontext"
	"github.com/bmviniciuss/cccat12/internal/ports"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type DriverHandler struct {
	createDriver *usecase.CreateDriver
	getDriver    *usecase.GetDriver
}

func NewDriverHandler(
	createDriver *usecase.CreateDriver,
	getDriver *usecase.GetDriver,
) *DriverHandler {
	return &DriverHandler{
		createDriver: createDriver,
		getDriver:    getDriver,
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
		return
	}
	out := &presentation.CreateDriverOutput{
		ID: res.ID,
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, out)
}

func (h *DriverHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqID, _ := customcontext.RequestID(ctx)
	id := chi.URLParam(r, "id")
	res, err := h.getDriver.Execute(ctx, id)
	if errors.Is(err, usecase.ErrDriverNotFound) {
		render.Render(w, r, presentation.ErrNotFound(reqID, err))
		return
	}
	if err != nil {
		render.Render(w, r, presentation.ErrUnprocessableEntity(reqID, err))
		return
	}
	render.Status(r, http.StatusOK)
	out := &presentation.GetDriverOutput{
		ID:          res.ID.String(),
		Name:        res.Name,
		Document:    res.Document.String(),
		PlateNumber: res.PlateNumber,
	}
	render.Render(w, r, out)
}
