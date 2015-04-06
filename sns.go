package main

import "github.com/awslabs/aws-sdk-go/service/sns"

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
