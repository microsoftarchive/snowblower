package aws

import (
	"github.com/awslabs/aws-sdk-go/service/sns"
)

type SNSPublisher struct {
	Service *sns.SNS
	Topic   string
}

func (p *SNSPublisher) Publish(message string) {
	input := sns.PublishInput{
		Message:  &message,
		TopicARN: &p.Topic,
	}
	p.Service.Publish(&input)
}
