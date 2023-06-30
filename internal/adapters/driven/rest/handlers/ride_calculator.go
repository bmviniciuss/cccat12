package handlers

import (
	"time"

	"github.com/bmviniciuss/cccat12/internal/adapters/driven/rest/presentation"
	"github.com/bmviniciuss/cccat12/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

type RideCalculatorHandler struct{}

func NewRideCalculatorHandler() *RideCalculatorHandler {
	return &RideCalculatorHandler{}
}

func (h *RideCalculatorHandler) Handle(c *fiber.Ctx) error {
	input := new(presentation.CalculateRideInput)
	if err := c.BodyParser(input); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	ride := entities.NewRide()
	for _, segment := range input.Segments {
		time, err := time.Parse(entities.TimeLayout, segment.Date)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": "Invalid Date",
			})
		}
		err = ride.AddSegment(segment.Distance, time)
		if err != nil {
			return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
				"message": err.Error(),
			})
		}
	}

	price := ride.Calculate()
	return c.JSON(presentation.CalculateRideOutput{
		Price: price,
	})
}
