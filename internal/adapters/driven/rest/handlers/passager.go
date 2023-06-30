package handlers

import (
	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/presentation"
	"github.com/bmviniciuss/cccat12/internal/application/usecase"
	"github.com/gofiber/fiber/v2"
)

type PassagerHandlerPort interface {
	Create(c *fiber.Ctx) error
}

type PassagerHandler struct {
	createPassager *usecase.CreatePassager
}

func NewPassagerHandler(createPassager *usecase.CreatePassager) *PassagerHandler {
	return &PassagerHandler{
		createPassager: createPassager,
	}
}

var (
	_ PassagerHandlerPort = (*PassagerHandler)(nil)
)

func (h *PassagerHandler) Create(c *fiber.Ctx) error {
	input := new(presentation.CreatePassagerInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	out, err := h.createPassager.Execute(c.Context(), usecase.CreatePassagerInput{
		Name:     input.Name,
		Email:    input.Email,
		Document: input.Document,
	})

	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res := presentation.CreatePassagerOutput{
		ID: out.ID,
	}
	return c.JSON(res)
}
