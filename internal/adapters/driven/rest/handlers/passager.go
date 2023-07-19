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

type PassengerHandler struct {
	createPassenger *usecase.CreatePassenger
	getPassenger    *usecase.GetPassenger
}

func NewPassengerHandler(
	createPassenger *usecase.CreatePassenger,
	getPassenger *usecase.GetPassenger,
) *PassengerHandler {
	return &PassengerHandler{
		createPassenger: createPassenger,
		getPassenger:    getPassenger,
	}
}

var (
	_ ports.PassengerHandlersPort = (*PassengerHandler)(nil)
)

func (h *PassengerHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqID, _ := customcontext.RequestID(ctx)

	var input presentation.CreatePassengerInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		render.Render(w, r, presentation.ErrBadRequest(reqID, err))
		return
	}

	out, err := h.createPassenger.Execute(ctx, usecase.CreatePassengerInput{
		Name:     input.Name,
		Email:    input.Email,
		Document: input.Document,
	})
	if err != nil {
		render.Render(w, r, presentation.ErrUnprocessableEntity(reqID, err))
		return
	}

	res := &presentation.CreatePassengerOutput{
		ID: out.ID,
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, res)
}

func (h *PassengerHandler) Get(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqID, _ := customcontext.RequestID(ctx)
	id := chi.URLParam(r, "id")

	res, err := h.getPassenger.Execute(ctx, id)
	if errors.Is(err, usecase.ErrorPassengerNotFound) {
		render.Render(w, r, presentation.ErrNotFound(reqID, err))
		return
	}
	if err != nil {
		render.Render(w, r, presentation.ErrUnprocessableEntity(reqID, err))
		return
	}

	out := &presentation.GetPassengerOutput{
		ID:    res.ID.String(),
		Name:  res.Name,
		Email: res.Email.Value,
	}
	render.Status(r, http.StatusOK)
	render.Render(w, r, out)
}
