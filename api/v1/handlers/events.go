package handlers

import (
	"event-messaging-api/internal/aws"
	"event-messaging-api/model"

	"github.com/gofiber/fiber/v2"
)

func SendEventToSNSTopic(c *fiber.Ctx) error {

	event_payment := new(model.Payment)

	if err := c.BodyParser(event_payment); err != nil {
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Payment Event!",
		})
	}

	go func () {
		cfg := aws.NewConfig()
		snsClient := aws.NewSNSClient(cfg)
		snsClient.PublishEvent(*event_payment)
	}()

	return c.JSON(fiber.Map{
		"code":    202,
		"message": "Created!",
	})
}
