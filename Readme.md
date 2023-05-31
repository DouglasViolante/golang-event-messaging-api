
# event-messaging-api

Brief project written with Go exclusively for educational purposes on the use of AWS SDK, GoFiber and OpenTelemetry. 

It simulates with AWS Localstack a publishing of an event in a SNS Topic when receiving an API Call.

## Third Party Libs and Tools

- [AWS SDK v2](https://aws.github.io/aws-sdk-go-v2/docs/)
- [GoFiber](https://docs.gofiber.io/)
- [Localstack](https://docs.localstack.cloud)
- [OpenTelemetry](https://opentelemetry.io/)
 
## Todo

- Telemetry (Tracing and Metrics)
- Fix DockerFile

## Observability

- [ZipKin](https://zipkin.io/)
- [Prometheus](https://prometheus.io/)

## How to Run

- Requires Docker!

- Run the command ```docker-compose up``` in the same folder level as the file ```docker-compose.yml```
