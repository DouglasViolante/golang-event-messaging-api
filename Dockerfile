FROM golang:1.19.5

ADD . /go/src/event-messaging-api
WORKDIR /go/src/event-messaging-api

COPY go.mod go.sum ./
RUN go mod download
COPY *.go ./

RUN go build -o /go/src/event-messaging-api/cmd/

EXPOSE 9101

CMD ["/go/src/event-messaging-api"]