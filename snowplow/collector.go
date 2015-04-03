package snowplow

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"

	"github.com/awslabs/aws-sdk-go/aws"
	"github.com/awslabs/aws-sdk-go/service/sns"
	"github.com/wunderlist/snowblower/snowplow"
)

var snsService *sns.SNS
var snsTopic string

const trackingPostPath = "/com.snowplowanalytics.snowplow/tp2"

func healthHandler(w http.ResponseWriter, request *http.Request) {
	msg := make(map[string]string)
	msg["up"] = "true"
	output, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}

func trackingPostHandler(w http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "POST":
		body, err := ioutil.ReadAll(request.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		} else if len(body) == 0 {
			log.Printf(
				"Empty body in request from %s %s",
				request.UserAgent(), request.Header.Get("X-Forwarded-For"))
		} else {
			payload := snowplow.Payload{}
			if err := json.Unmarshal(body, &payload); err != nil {
				log.Printf("Rejecting posted payload. Error %s Payload: %s", err, body)
				w.WriteHeader(http.StatusBadRequest)
			} else {
				if err := processPayload(payload); err != nil {
					w.WriteHeader(http.StatusServiceUnavailable)
				} else {
					w.WriteHeader(http.StatusOK)
				}
			}
		}
	default:
		w.WriteHeader(http.StatusForbidden)
	}
	w.Write(nil)
}

func processPayload(payload snowplow.Payload) error {
	for _, event := range payload.Data {

		msg := string(event)
		input := sns.PublishInput{
			Message:  &msg,
			TopicARN: &snsTopic,
		}
		snsService.Publish(&input)

	}
	return nil
}

// RunCollector starts the HTTP collector
func RunCollector() {
	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	var credentials aws.CredentialsProvider
	if os.Getenv("AWS_ACCESS_KEY_ID") != "" {
		credentials = aws.DefaultCreds()
	} else {
		credentials = aws.IAMCreds()
	}

	snsService = sns.New(&aws.Config{
		Credentials: credentials,
		Region:      "eu-west-1"})

	snsTopic = os.Getenv("SNS_TOPIC")
	if snsTopic == "" {
		panic("SNS_TOPIC required")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	http.HandleFunc("/api/health", healthHandler)
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc(trackingPostPath, trackingPostHandler)
	http.ListenAndServe(fmt.Sprintf(":%s", port), nil)
}
