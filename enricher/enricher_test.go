package enricher

import (
	"fmt"
	"testing"
	"encoding/json"
	
	"github.com/wunderlist/snowblower/common"
	sp "github.com/wunderlist/snowblower/snowplow"

	"github.com/awslabs/aws-sdk-go/service/sqs"
)

const simpleMessage = `
{"Type":"Notification","MessageId":"18fd4a34-443a-5289-971a-07f40247f98f","TopicArn":"arn:aws:sns:eu-west-1:030398754360:snowblower","Message":"{\"schema\":\"iglu:com.snowplowanalytics.snowplow/CollectorPayload/thrift/1-0-0\",\"ipAddress\":\"181.189.182.158\",\"timestamp\":1430668349024,\"encoding\":\"\",\"collector\":\"Snowblower/0.0.1\",\"userAgent\":\"Wunderlist/21 CFNetwork/711.3.18 Darwin/14.0.0\",\"body\":\"{\\\"schema\\\":\\\"iglu:com.snowplowanalytics.snowplow\\\\/payload_data\\\\/jsonschema\\\\/1-0-0\\\",\\\"data\\\":[{\\\"aid\\\":\\\"Wunderlist\\\\/3.2.3 9f43764ad52624fcb734633947fcfb273798111b\\\",\\\"res\\\":\\\"640x1136\\\",\\\"ue_px\\\":\\\"eyJzY2hlbWEiOiJpZ2x1OmNvbS5zbm93cGxvd2FuYWx5dGljcy5zbm93cGxvd1wvdW5zdHJ1Y3RfZXZlbnRcL2pzb25zY2hlbWFcLzEtMC0wIiwiZGF0YSI6eyJzY2hlbWEiOiJpZ2x1OmNvbS53dW5kZXJsaXN0XC9jbGllbnRfZXZlbnRcL2pzb25zY2hlbWFcLzEtMC0wIiwiZGF0YSI6eyJldmVudCI6ImNsaWVudC5zeW5jLm1hdHJ5b3Noa2EiLCJwYXJhbWV0ZXJzIjp7ImR1cmF0aW9uIjoyMTI1Ljg2Mjk1NjA0NzA1OCwic3VjY2VzcyI6dHJ1ZX19fX0\\\",\\\"uid\\\":\\\"14119542\\\",\\\"p\\\":\\\"mob\\\",\\\"cx\\\":\\\"eyJzY2hlbWEiOiJpZ2x1OmNvbS5zbm93cGxvd2FuYWx5dGljcy5zbm93cGxvd1wvY29udGV4dHNcL2pzb25zY2hlbWFcLzEtMC0wIiwiZGF0YSI6W3sic2NoZW1hIjoiaWdsdTpjb20uc25vd3Bsb3dhbmFseXRpY3Muc25vd3Bsb3dcL21vYmlsZV9jb250ZXh0XC9qc29uc2NoZW1hXC8xLTAtMSIsImRhdGEiOnsib3NUeXBlIjoiaW9zIiwiZGV2aWNlTWFudWZhY3R1cmVyIjoiQXBwbGUgSW5jLiIsIm9wZW5JZGZhIjoiREIyMERCMTEtM0IzQS0zMjNDLUE2N0YtOURGNTkxQTBGN0UzIiwiY2FycmllciI6IlRJR08iLCJkZXZpY2VNb2RlbCI6ImlQaG9uZSIsIm9zVmVyc2lvbiI6IjguMyIsIm5ldHdvcmtUeXBlIjoibW9iaWxlIiwiYXBwbGVJZGZ2IjoiMDJERUFCNUMtODVFOS00MTdDLUE1OUUtNzA4MTdCQUZBNjBEIiwibmV0d29ya1RlY2hub2xvZ3kiOiJDVFJhZGlvQWNjZXNzVGVjaG5vbG9neUhTRFBBIiwiYXBwbGVJZGZhIjoiNkQwNkFDQzMtNDE1OC00NzFCLTg1MDUtMURGNkNGRjQ3NDYzIn19XX0\\\",\\\"dtm\\\":\\\"1430668347186\\\",\\\"tv\\\":\\\"ios-0.3.2\\\",\\\"tna\\\":\\\"com.wunderlist\\\",\\\"lang\\\":\\\"es\\\",\\\"e\\\":\\\"ue\\\",\\\"vp\\\":\\\"640x1136\\\",\\\"eid\\\":\\\"d48b1ba7-4aa3-49d8-bcb8-649d475e615d\\\"}]}\",\"headers\":[\"Cookie: sp=98d126a7-0b4f-4a10-b863-a214536fedf5; sp=98d126a7-0b4f-4a10-b863-a214536fedf5\",\"Connection: keep-alive\",\"Accept-Encoding: gzip, deflate\",\"Accept-Language: es-es\",\"Content-Type: application/x-www-form-urlencoded\",\"X-Forwarded-For: 181.189.182.158\",\"X-Forwarded-Port: 443\",\"X-Forwarded-Proto: https\",\"Content-Length: 1392\",\"Accept: text/html, application/x-www-form-urlencoded, text/plain, image/gif\",\"User-Agent: Wunderlist/21 CFNetwork/711.3.18 Darwin/14.0.0\"],\"contentType\":\"\",\"hostname\":\"\",\"networkUserId\":\"98d126a7-0b4f-4a10-b863-a214536fedf5\"}","Timestamp":"2015-05-03T15:52:29.034Z","SignatureVersion":"1","Signature":"gTP6p6pQpVrShbrhVik7pFbjsv/tjhWtzHfZOCyZixMM0FGJ1Pd+sQLueP1X5SlwDPbl7DtYZJ00eo4dGrh57cYDWeZJfDk2eKcgl8WrsoSqiDXHcY95xlktCBn2CsQjdC8dOK3muhioD2ad8Y94gXPlbTxoj1241xKE5XSnk1U2+PMFL77cMMa+xhbZ+tLz2bMu6jxvw161pS3VwHekUApJaQ4SyyrSBYHP5bUw5qgfY7OAru0lamZZAIrr07i7M84uEbFgJ9NO3JegoQzF1pztoDlTgI/5v6rYmxFkp+6gbDgQECKdiGS4P2+Uz80hIlSYpxDXfTiKVv7K9kMnCQ==","SigningCertURL":"https://sns.eu-west-1.amazonaws.com/SimpleNotificationService-d6d679a1d18e95c2f9ffcf11f4f9e198.pem","UnsubscribeURL":"https://sns.eu-west-1.amazonaws.com/?Action=Unsubscribe&SubscriptionArn=arn:aws:sns:eu-west-1:030398754360:snowblower:61d04990-a98a-46af-885a-3949a059beaa"}
`

type testPublisher struct {
	events []string
}

func (p *testPublisher) Publish(message string) {
	p.events = append(p.events, message)
}

func setupSqsMessage(body string) *sqs.Message {
	return &sqs.Message{Body: &body}
}

func setupTestPublisher() common.Publisher {
	var publisher common.Publisher
	publisher = &testPublisher{}
	return publisher
}

func extractEventsFromTestPublisher(publisher *testPublisher, t *testing.T) []sp.Event {
	var events []sp.Event

	for _, e := range publisher.events {
		var event sp.Event
		if err := json.Unmarshal([]byte(e), &event); err != nil {
			t.Errorf("Failed to unmarshal a created event: %s\n", err)
		} else {
			events = append(events, event)
		}
	}

	return events
}

func TestProcessSimpleMessage(t *testing.T) {
	message := setupSqsMessage(simpleMessage)
	publisher := setupTestPublisher()

	processSNSMessage(message, publisher)
	events := extractEventsFromTestPublisher(publisher.(*testPublisher), t)

	if len(events) != 1 {
		t.Errorf("Shit, got %v events", len(events))
	}

	event := events[0]

	// redundant?
	if event.AppID != "Wunderlist/3.2.3 9f43764ad52624fcb734633947fcfb273798111b" {
		t.Errorf("Shit, AppID field is incorrect!\n\thave: %v\n\texpect: %v\n", event.AppID, "Wunderlist/3.2.3 9f43764ad52624fcb734633947fcfb273798111b")
	}
	
	// redundant?
	if event.Platform != "mob" {
		t.Errorf("Shit, Platform field is incorrect!\n\thave: %v\n\texpect: %v\n", event.Platform, "mob")
	}

	// TODO: test ETLTimestamp
	
	if event.CollectorTimestamp != "1430668349024" {
		t.Errorf("Shit, CollectorTimestamp field is incorrect!\n\thave: %v\n\texpect: %v\n", event.CollectorTimestamp, "1430668349024")
	}

	// redundant?
	if event.DeviceTimestamp != "1430668347186" {
		t.Errorf("Shit, DeviceTimestamp field is incorrect!\n\thave: %v\n\texpect: %v\n", event.DeviceTimestamp, "1430668347186")
	}
	


	//fmt.Println(storer.(testStorer).events)
	fmt.Printf("length of events: %d", len(publisher.(*testPublisher).events))
}
