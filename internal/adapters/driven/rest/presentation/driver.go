package presentation

import "net/http"

type CreateDriverInput struct {
	Name        string `json:"name"`
	Document    string `json:"document"`
	PlateNumber string `json:"plate_number"`
}

type CreateDriverOutput struct {
	ID string `json:"id"`
}

func (cdo *CreateDriverOutput) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

type GetDriverOutput struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Document    string `json:"document"`
	PlateNumber string `json:"plate_number"`
}

func (gdo *GetDriverOutput) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}
