package aws

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

type Config struct {
	session           *session.Session
}

func NewConfig() *Config {
	sess, err := session.NewSessionWithOptions(session.Options{
		Profile: "localstack",
		Config: aws.Config{
			Region: aws.String("us-east-1"),
			Endpoint: aws.String("http://localhost:4566"),
			CredentialsChainVerboseErrors: aws.Bool(true),
		},
	})

	if err != nil {
		log.Fatalln(err)
	}

	return &Config{
		session:           sess,
	}
}