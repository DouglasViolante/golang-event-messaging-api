package router

import (
	"event-messaging-api/api/v1/handlers"
	"event-messaging-api/middleware"

	"github.com/gofiber/fiber/v2"
)

func Initalize(router *fiber.App) {
	router.Use(middleware.Json)

	events := router.Group("/events/payment")

	events.Post("/newPayment/", handlers.SendEventToSNSTopic)

	router.Use(func(c *fiber.Ctx) error {
		return c.Status(404).JSON(fiber.Map{
			"code":    404,
			"message": "404: Not Found",
		})
	})
}
