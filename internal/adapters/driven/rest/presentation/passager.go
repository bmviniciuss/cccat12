package presentation

import "net/http"

type CreatePassengerInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
}

type CreatePassengerOutput struct {
	ID string `json:"id"`
}

func (o *CreatePassengerOutput) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
