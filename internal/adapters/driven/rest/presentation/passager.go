package presentation

type CreatePassagerInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
}

type CreatePassagerOutput struct {
	ID string `json:"id"`
}
