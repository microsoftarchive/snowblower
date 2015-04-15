package main

import (
	"fmt"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/sns"
	"github.com/awslabs/aws-sdk-go/service/sqs"
)

// SNSPublisher ...
type SNSPublisher struct {
	service *sns.SNS
	topic   string
}

func (p *SNSPublisher) publish(message string) {
	input := sns.PublishInput{
		Message:  &message,
		TopicARN: &p.topic,
	}
	p.service.Publish(&input)
}

// SQSService ...
type SQSService struct {
	service *sqs.SQS
	url     string
}

func (service *SQSService) getMessages() ([]*sqs.Message, error) {
	params := &sqs.ReceiveMessageInput{
		QueueURL: aws.String(service.url),
		AttributeNames: []*string{
			aws.String("SentTimestamp"),
		},
		MaxNumberOfMessages: aws.Long(10),
		MessageAttributeNames: []*string{
			aws.String("All"),
		},
		VisibilityTimeout: aws.Long(60),
		WaitTimeSeconds:   aws.Long(20),
	}
	response, err := service.service.ReceiveMessage(params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	if response != nil {
		return nil, err
	}
	return response.Messages, nil
}
