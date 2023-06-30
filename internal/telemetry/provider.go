package telemetry

import (
	"context"
	"log"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func InitProvider() func() {
	ctx := context.Background()

	res, err := resource.New(ctx,
		resource.WithProcess(),
		resource.WithFromEnv(),
		resource.WithHost(),
		resource.WithTelemetrySDK(),
		resource.WithAttributes(
			semconv.ServiceNameKey.String("event-messaging-api"),
			semconv.TelemetrySDKLanguageGo,
		),
	)
	handleError(err, " ERROR | Error at Resource Creation for Telemetry Backend! ")

	otelAgentAddr := "localhost:4317" 				// For Local Debugging Only!

	//Trace Configuration
	traceClient := otlptracegrpc.NewClient(
		otlptracegrpc.WithEndpoint(otelAgentAddr),	// For Local Debugging Only!
		otlptracegrpc.WithInsecure(), 				// For Local Debugging Only!
		//otlptracegrpc.WithDialOption(grpc.WithBlock()),
	)
	traceExporter, err := otlptrace.New(ctx, traceClient)
	handleError(err, " ERROR | Error at Trace Collector Creation! ")

	bsp := sdktrace.NewBatchSpanProcessor(traceExporter)
	traceProvider := sdktrace.NewTracerProvider(
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithSpanProcessor(bsp),
	)
	otel.SetTracerProvider(traceProvider)

	//Metric Configuration
	metricExporter, err := otlpmetricgrpc.New(
		ctx,
		otlpmetricgrpc.WithEndpoint(otelAgentAddr),	// For Local Debugging Only!
		otlpmetricgrpc.WithInsecure(),				// For Local Debugging Only!
	)
	handleError(err, " ERROR | Error at Metric Collector Creation! ")
	
	meterProvider := sdkmetric.NewMeterProvider(
		sdkmetric.WithResource(res),
		sdkmetric.WithReader(
			sdkmetric.NewPeriodicReader(
				metricExporter,
				sdkmetric.WithInterval(2*time.Second),
			),
		),
	)
	otel.SetMeterProvider(meterProvider)

	//Setting Global Propagator - NO-OP Default
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	//otel.SetTracerProvider(traceProvider)

	return func() {
		cxt, cancel := context.WithTimeout(ctx, time.Second)
		defer cancel()
		if err := traceExporter.Shutdown(cxt); err != nil{
			otel.Handle(err)
		}
		if err := meterProvider.Shutdown(cxt); err != nil{
			otel.Handle(err)
		}
	}

}

func handleError(err error, message string) {
	if err != nil {
		log.Fatalf("%s: %v", message, err)
	}
}