package telemetry

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var metricInstance metric.Meter

func GetMeter() metric.Meter {
	if metricInstance == nil {
		metricInstance = otel.Meter("event-messaging-api")
	}
	return metricInstance
}