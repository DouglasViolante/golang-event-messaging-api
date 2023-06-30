package aws

import (
	"encoding/json"
	"event-messaging-api/model"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go/service/sns"
)

type SNSClient struct {
	svc *sns.SNS
	cfg *Config
}

func NewSNSClient(c *Config) *SNSClient {
	return &SNSClient{
		sns.New(c.session),
		c,
	}
}

func (s SNSClient) PublishEvent(payment model.Payment) {

    messageBytes, _ := json.Marshal(payment)
    messageStr := string(messageBytes)

    req := &sns.PublishInput{
        TopicArn: aws.String("arn:aws:sns:us-east-1:000000000000:payments-topic"),
        Message: aws.String(messageStr),}

    res, err := s.svc.Publish(req)
	
    if err != nil {
		log.Println(err.Error())
    } else {
		fmt.Printf("PROCESSING FINISHED! %s", *res.MessageId)
	}
}