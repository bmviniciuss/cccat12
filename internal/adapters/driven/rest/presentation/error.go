package presentation

import (
	"net/http"

	"github.com/go-chi/render"
)

type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	ID      string `json:"id"`
	Message string `json:"message"`
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}

func ErrRender(id string, err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 500,
		ID:             id,
		Message:        "Internal Server Error",
	}
}

func ErrBadRequest(id string, err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 400,
		ID:             id,
		Message:        "Bad Request",
	}
}

func ErrNotFound(id string, err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 404,
		ID:             id,
		Message:        "Not Found",
	}
}

func ErrUnprocessableEntity(id string, err error) *ErrResponse {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		ID:             id,
		Message:        err.Error(),
	}
}
