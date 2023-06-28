package handlers

import (
	"time"

	"github.com/bmviniciuss/cccat12/internal/domain/entities"
	"github.com/gofiber/fiber/v2"
)

type CalculateRideInput struct {
	Segments []Segment `json:"segments"`
}

type Segment struct {
	Distance float64 `json:"distance"`
	Date     string  `json:"date"`
}

func CalculateRideHandler(c *fiber.Ctx) error {
	input := new(CalculateRideInput)
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
				"message": err.Error(),
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
	return c.JSON(fiber.Map{
		"price": price,
	})

}
