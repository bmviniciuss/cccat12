package handlers

import (
	"github.com/bmviniciuss/cccat12/internal/application/usecase"
	"github.com/gofiber/fiber/v2"
)

type CreatePassagerInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Document string `json:"document"`
}

type CreatePassagerOutput struct {
	ID string `json:"id"`
}

func CreatePassagerHandler(c *fiber.Ctx) error {
	input := new(CreatePassagerInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	uc := usecase.NewCreatePassager()
	out, err := uc.Execute(c.Context(), usecase.CreatePassagerInput{
		Name:     input.Name,
		Email:    input.Email,
		Document: input.Document,
	})

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res := CreatePassagerOutput{
		ID: out.ID,
	}
	return c.JSON(res)
}
