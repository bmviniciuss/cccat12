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

type PassagerHandler struct {
	createPassager *usecase.CreatePassager
}

func NewPassagerHandler(createPassager *usecase.CreatePassager) *PassagerHandler {
	return &PassagerHandler{
		createPassager: createPassager,
	}
}

var (
	_ ports.PassagerHandlersPort = (*PassagerHandler)(nil)
)

func (h *PassagerHandler) Create(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	reqID, _ := customcontext.RequestID(ctx)

	var input presentation.CreatePassagerInput
	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		render.Render(w, r, presentation.ErrBadRequest(reqID, err))
		return
	}

	out, err := h.createPassager.Execute(ctx, usecase.CreatePassagerInput{
		Name:     input.Name,
		Email:    input.Email,
		Document: input.Document,
	})
	if err != nil {
		render.Render(w, r, presentation.ErrUnprocessableEntity(reqID, err))
		return
	}

	res := &presentation.CreatePassagerOutput{
		ID: out.ID,
	}
	render.Status(r, http.StatusCreated)
	render.Render(w, r, res)
}
