package presentation

type CreateDriverInput struct {
	Name        string `json:"name"`
	Document    string `json:"document"`
	PlateNumber string `json:"plate_number"`
}

type CreateDriverOutput struct {
	ID string `json:"id"`
}
