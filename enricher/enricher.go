package enricher

import (
	"encoding/json"
	"fmt"

	"github.com/wunderlist/snowblower/common"
	sp "github.com/wunderlist/snowblower/snowplow"
	sb_aws "github.com/wunderlist/snowblower/aws"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/sqs"
	"github.com/duncan/base64x"
)

var queue struct {
	service *sqs.SQS
	params  *sqs.ReceiveMessageInput
}

type enrichedPublisher struct {
	events []sp.Event
}

func (s *enrichedPublisher) publish(e *sp.Event) {
	s.events = append(s.events, *e)
}

var publisher Publisher

func Start(config common.Config) {
	queue.service = sqs.New(&aws.Config{
		Credentials: config.Credentials,
		Region:      "eu-west-1",
	})

	queue.params = &sqs.ReceiveMessageInput{
		QueueURL: aws.String(config.CollectedSqsURL),
		AttributeNames: []*string{
			aws.String("All"), // Required
		},
		MaxNumberOfMessages: aws.Long(1),
		MessageAttributeNames: []*string{
			aws.String("All"), // Required
		},
		VisibilityTimeout: aws.Long(3600),
		WaitTimeSeconds:   aws.Long(10),
	}

	publisher = &enrichedPublisher{}

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

	for _, message := range resp.Messages {
		processSNSMessage(message, publisher)
	}
}

func processSNSMessage(message *sqs.Message, p Publisher) {
	//messageID := *message.MessageID
	//receiptHandle := *message.ReceiptHandle

	snsMessage := sb_aws.SNSMessage{}
	if err := json.Unmarshal([]byte(*message.Body), &snsMessage); err != nil {
		fmt.Printf("SNS MESSAGE UNMARSHALL ERROR %s\n", err)
	} else {
		payload := sp.CollectorPayload{}
		if err := json.Unmarshal([]byte(snsMessage.Message), &payload); err != nil {
			fmt.Printf("COLLECTOR PAYLOAD UNMARSHALL ERROR %s\n", err)
		} else {
			processCollectorPayload(payload, p)
			// schedule for deletion
		}
	}
}

func processCollectorPayload(cp sp.CollectorPayload, p Publisher) {
	tp := sp.TrackerPayload{}
	if err := json.Unmarshal([]byte(cp.Body), &tp); err != nil {
		fmt.Printf("TRACKER PAYLOAD UNMARSHALL ERROR %s\n", err)
	} else {
		for _, e := range tp.Data {
			//dsfmt.Printf("%s, %s", cp.NetworkUserID, e.AppID)
			processEvent(e, tp, cp, p)
		}
	}
}

func processEvent(e sp.Event, tp sp.TrackerPayload, cp sp.CollectorPayload, p Publisher) {
	enrichEvent(&e, tp, cp)
	publishEvent(&e, p)
	
	o, _ := json.MarshalIndent(e, "", " ")
	fmt.Printf("JSON: %s", o)
}

func enrichEvent(e *sp.Event, tp sp.TrackerPayload, cp sp.CollectorPayload) {
	b, _ := base64x.URLEncoding.DecodeString(e.UnstructuredEventEncoded)
	ue := sp.Iglu{}
	if err := json.Unmarshal(b, &ue); err != nil {
		fmt.Printf("UE UNMARSHALL ERROR %s\n", err)
		return
	}
	b, _ = json.Marshal(ue)
	e.UnstructuredEvent = string(b)

	b, _ = base64x.URLEncoding.DecodeString(e.ContextsEncoded)
	co := sp.Iglu{}
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
}

func publishEvent(e *sp.Event, p Publisher) {
	p.publish(e)
	//fmt.Printf("\nXXX length of events: %d\n", len(p.(*testPublisher).events))
}
