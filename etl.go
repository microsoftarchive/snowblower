package main

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/sqs"
	"github.com/duncan/base64x"
)

var queue struct {
	service *sqs.SQS
	params  *sqs.ReceiveMessageInput
}

func startETL() {
	queue.service = sqs.New(config.awsSession, &aws.Config{Region: aws.String(config.awsregion)})

	queue.params = &sqs.ReceiveMessageInput{
		QueueUrl: aws.String(config.sqsURL),
		AttributeNames: []*string{
			aws.String("All"), // Required
		},
		MaxNumberOfMessages: aws.Int64(1),
		MessageAttributeNames: []*string{
			aws.String("All"), // Required
		},
		VisibilityTimeout: aws.Int64(3600),
		WaitTimeSeconds:   aws.Int64(10),
	}

	// while something....
	processNextBatch()
}

func processNextBatch() {

	resp, err := queue.service.ReceiveMessage(queue.params)

	if err != nil {
		// A service error occurred.
		fmt.Println("Error:", err.Error())
	}

	for _, message := range resp.Messages {
		processSNSMessage(message)
	}
}

func processSNSMessage(message *sqs.Message) {
	//messageID := *message.MessageID
	//receiptHandle := *message.ReceiptHandle

	snsMessage := SNSMessage{}

	if err := json.Unmarshal([]byte(*message.Body), &snsMessage); err != nil {
		fmt.Printf("SNS MESSAGE UNMARSHALL ERROR %s\n", err)
	} else {
		payload := CollectorPayload{}
		if err := json.Unmarshal([]byte(snsMessage.Message), &payload); err != nil {
			fmt.Printf("COLLECTOR PAYLOAD UNMARSHALL ERROR %s\n", err)
		} else {
			processCollectorPayload(payload)
			// schedule for deletion
		}
	}
}

func processCollectorPayload(cp CollectorPayload) {
	tp := TrackerPayload{}
	if err := json.Unmarshal([]byte(cp.Body), &tp); err != nil {
		fmt.Printf("TRACKER PAYLOAD UNMARSHALL ERROR %s\n", err)
	} else {
		for _, e := range tp.Data {
			//dsfmt.Printf("%s, %s", cp.NetworkUserID, e.AppID)
			processEvent(e, tp, cp)
		}
	}
}

func processEvent(e Event, tp TrackerPayload, cp CollectorPayload) {
	b, _ := base64x.URLEncoding.DecodeString(e.UnstructuredEventEncoded)
	ue := Iglu{}
	if err := json.Unmarshal(b, &ue); err != nil {
		fmt.Printf("UE UNMARSHALL ERROR %s\n", err)
		return
	}
	b, _ = json.Marshal(ue)
	e.UnstructuredEvent = string(b)

	b, _ = base64x.URLEncoding.DecodeString(e.ContextsEncoded)
	co := Iglu{}
	if err := json.Unmarshal(b, &co); err != nil {
		fmt.Printf("CO UNMARSHALL ERROR %s\n", err)
		return
	}
	b, _ = json.Marshal(co)
	e.Contexts = string(b)

	// pick up details from colletor payload
	e.UserIPAddress = cp.IPAddress
	e.CollectorTimestamp = string(cp.Timestamp)
	e.CollectorVersion = cp.Collector
	e.UserAgent = cp.UserAgent
	// cp.RefererURI
	e.PageURLPath = cp.Path
	e.PageURLQuery = cp.QueryString
	// cp.Headers
	e.NetworkUserID = cp.NetworkUserID

	o, _ := json.MarshalIndent(e, "", " ")
	fmt.Printf("JSON: %s", o)

}
