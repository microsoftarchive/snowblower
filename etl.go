package main

import (
	"fmt"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/aws/awsutil"
	"github.com/awslabs/aws-sdk-go/service/sqs"
)

var queue struct {
	service *sqs.SQS
	params  *sqs.ReceiveMessageInput
}

func startETL() {
	queue.service = sqs.New(&aws.Config{
		Credentials: config.credentials,
		Region:      "eu-west-1",
	})

	queue.params = &sqs.ReceiveMessageInput{
		QueueURL: aws.String(config.sqsURL),
		AttributeNames: []*string{
			aws.String("SentTimestamp"), // Required
		},
		MaxNumberOfMessages: aws.Long(10),
		MessageAttributeNames: []*string{
			aws.String("All"), // Required
		},
		VisibilityTimeout: aws.Long(3600),
		WaitTimeSeconds:   aws.Long(10),
	}

	// while something....
	processNextBatch()
}

func processNextBatch() {

	resp, err := queue.service.ReceiveMessage(queue.params)

	if awserr := aws.Error(err); awserr != nil {
		// A service error occurred.
		fmt.Println("Error:", awserr.Code, awserr.Message)
	} else if err != nil {
		// A non-service error occurred.
		panic(err)
	}

	fmt.Println(awsutil.StringValue(resp))

}
