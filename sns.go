package main

import (
	"encoding/json"
	"log"

	"github.com/awslabs/aws-sdk-go/service/sns"
)

type snsPublisher struct {
	service *sns.SNS
	topic   string
}

func (p *snsPublisher) publish(event *Event) {
	bytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		return
	}
	message := string(bytes)
	input := sns.PublishInput{
		Message:  &message,
		TopicARN: &p.topic,
	}
	p.service.Publish(&input)
}
