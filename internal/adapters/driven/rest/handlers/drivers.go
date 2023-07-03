package handlers

import (
	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/presentation"
	"github.com/bmviniciuss/cccat12/internal/application/usecase"
	"github.com/gofiber/fiber/v2"
)

type DriverHandlerPort interface {
	Create(c *fiber.Ctx) error
}

type DriverHandler struct {
	createDriver *usecase.CreateDriver
}

func NewDriverHandler(createDriver *usecase.CreateDriver) *DriverHandler {
	return &DriverHandler{
		createDriver: createDriver,
	}
}

var (
	_ DriverHandlerPort = (*DriverHandler)(nil)
)

func (h *DriverHandler) Create(c *fiber.Ctx) error {
	input := new(presentation.CreateDriverInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	res, err := h.createDriver.Execute(c.Context(), usecase.CreateDriverInput{
		Name:        input.Name,
		Document:    input.Document,
		PlateNumber: input.PlateNumber,
	})
	if err != nil {
		return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	out := presentation.CreateDriverOutput{
		ID: res.ID,
	}

	return c.JSON(out)
}
