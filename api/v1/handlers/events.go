package handlers

import (
	"context"
	"event-messaging-api/internal/aws"
	"event-messaging-api/internal/telemetry"
	"event-messaging-api/model"

	"github.com/gofiber/fiber/v2"
	"go.opentelemetry.io/otel/codes"
)

func SendEventToSNSTopic(c *fiber.Ctx) error {

	ctx := context.TODO()

	_, span := telemetry.GetTracer().Start(ctx, "handler::SendEventToSNSTopic")
	defer span.End()
	eventCounter, _ := telemetry.GetMeter().Int64Counter("counter.total.sns.events")

	event_payment := new(model.Payment)

	if err := c.BodyParser(event_payment); err != nil {
		span.SetStatus(codes.Error, "Invalid Payment Event!")
		span.RecordError(err)
		return c.JSON(fiber.Map{
			"code":    400,
			"message": "Invalid Payment Event!",
		})
	}

	go func () {
		cfg := aws.NewConfig()
		snsClient := aws.NewSNSClient(cfg)
		snsClient.PublishEvent(*event_payment)
		eventCounter.Add(ctx, 1)
	}()

	return c.JSON(fiber.Map{
		"code":    202,
		"message": "Created!",
	})
}
