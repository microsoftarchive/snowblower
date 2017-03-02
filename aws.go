package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sns"
	"github.com/aws/aws-sdk-go/service/sqs"
)

// SNSMessage ...
type SNSMessage struct {
	Type      string `json:"Type"`
	MessageID string `json:"MessageID"`
	Message   string `json:"Message"`
}

// SNSPublisher ...
type SNSPublisher struct {
	service *sns.SNS
	topic   string
}

func (p *SNSPublisher) publish(message string) {
	input := sns.PublishInput{
		Message:  &message,
		TopicArn: &p.topic,
	}

	resp, err := p.service.Publish(&input)

	if err != nil {
		// Print the error, cast err to awserr.Error to get the Code and
		// Message from an error.
		fmt.Println(err.Error())
		return
	}

	// Pretty-print the response data.
	fmt.Println("Pushed: ", *resp.MessageId)
}

// SQSService ...
type SQSService struct {
	service *sqs.SQS
	url     string
}

func (service *SQSService) getMessages() ([]*sqs.Message, error) {
	params := &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(service.url),
		AttributeNames: []*string{
			aws.String("SentTimestamp"),
		},
		MaxNumberOfMessages: aws.Int64(10),
		MessageAttributeNames: []*string{
			aws.String("All"),
		},
		VisibilityTimeout: aws.Int64(60),
		WaitTimeSeconds:   aws.Int64(20),
	}
	response, err := service.service.ReceiveMessage(params)

	if err != nil {
		// A service error occurred.
		fmt.Println("Error:", err.Error())
	}

	if response != nil {
		return nil, err
	}
	return response.Messages, nil
}
