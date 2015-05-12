package enricher

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

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
	events []string
}

func (s *enrichedPublisher) Publish(message string) {
	s.events = append(s.events, message)
}

var publisher common.Publisher

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

func processSNSMessage(message *sqs.Message, p common.Publisher) {
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

func processCollectorPayload(cp sp.CollectorPayload, p common.Publisher) {
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

func processEvent(e sp.Event, tp sp.TrackerPayload, cp sp.CollectorPayload, p common.Publisher) {
	enrichEvent(&e, tp, cp)
	publishEvent(&e, p)

	cpo, _ := json.MarshalIndent(cp, "", " ")
	fmt.Printf("\n\nCollectorPayload:\n%s\n", cpo)

	tpo, _ := json.MarshalIndent(tp, "", " ")
	fmt.Printf("\n\nTrackerPayload:\n%s\n", tpo)

	o, _ := json.MarshalIndent(e, "", " ")
	fmt.Printf("\n\nEvent: \n%s\n\n", o)
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

	o, _ := json.MarshalIndent(ue, "", " ")
	fmt.Printf("\nUnstructuredEvent: \n%s\n", o)

	b, _ = base64x.URLEncoding.DecodeString(e.ContextsEncoded)
	co := sp.Iglu{}
	if err := json.Unmarshal(b, &co); err != nil {
		fmt.Printf("CO UNMARSHALL ERROR %s\n", err)
		return
	}
	b, _ = json.Marshal(co)
	e.Contexts = string(b)

	o, _ = json.MarshalIndent(co, "", " ")
	fmt.Printf("\nContexts: \n%s\n", o)

	// Set ETL timestamp
	e.ETLTimestamp = strconv.FormatInt(time.Now().UnixNano() / 1000000, 10)

	// pick up details from colletor payload
	e.UserIPAddress = cp.IPAddress
	e.CollectorTimestamp = strconv.FormatInt(cp.Timestamp, 10)
	e.CollectorVersion = cp.Collector
	e.UserAgent = cp.UserAgent
	// cp.RefererURI
	e.PageURLPath = cp.Path
	e.PageURLQuery = cp.QueryString
	// cp.Headers
	e.NetworkUserID = cp.NetworkUserID

	// from context
	contextData := co.Data.([]interface{})[0].(map[string]interface{})["data"].(map[string]interface{})
	e.DeviceType = contextData["deviceModel"].(string)
	e.OSFamily = contextData["osType"].(string)
	e.DeviceIsMobile = isMobile(e.OSFamily)
	fmt.Printf("%v", contextData["deviceModel"])

	//e.OSFamily = 
}

func isMobile(osType string) bool {
	switch osType {
		case "android", "ios": return true
		default: return false
	}
}

func publishEvent(e *sp.Event, p common.Publisher) {
	// TODO: serialise event.
	s, _ := json.Marshal(e)
	p.Publish(string(s))
	//fmt.Printf("\nXXX length of events: %d\n", len(p.(*testPublisher).events))
}
