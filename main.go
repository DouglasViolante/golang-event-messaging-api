package main

import (
	"event-messaging-api/api/v1/router"
	"event-messaging-api/internal/telemetry"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	shutdown := telemetry.InitProvider()
	defer shutdown()

	app := fiber.New(fiber.Config{AppName: "event-messaging-api"})

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	router.Initalize(app)
	log.Fatal(app.Listen(":" + "9101"))
}
