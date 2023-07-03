package presentation

import "net/http"

type CreatePassagerInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
}

type CreatePassagerOutput struct {
	ID string `json:"id"`
}

func (o *CreatePassagerOutput) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
