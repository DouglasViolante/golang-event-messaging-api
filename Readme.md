
# event-messaging-api

A brief project written with Go exclusively for educational purposes on the use of AWS SDK, GoFiber and OpenTelemetry. 

It simulates AWS environment with AWS Localstack for publishing an event in a SNS Topic when receiving an API Call, and also generates simples metrics and traces. 

## Third Party Libs and Tools

- [AWS SDK v2](https://aws.github.io/aws-sdk-go-v2/docs/)
- [GoFiber](https://docs.gofiber.io/)
- [Localstack](https://docs.localstack.cloud)
- [OpenTelemetry](https://opentelemetry.io/)
 
## Todo

- Fix DockerFile

## Observability

- [ZipKin](https://zipkin.io/)
- [Prometheus](https://prometheus.io/)

## How to Run

- Requires Docker!

- Run the command ```docker-compose up --force-recreate --build``` in the same folder level as the file ```docker-compose.yml```
