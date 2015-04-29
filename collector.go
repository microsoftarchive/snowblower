package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/wunderlist/snowblower/snowplow"

	"code.google.com/p/go-uuid/uuid"
)

type collector struct {
	publisher Publisher
}

func (c *collector) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	var networkID string
	spCookie, err := request.Cookie("sp")
	if err != nil {
		networkID = uuid.NewRandom().String()
	} else {
		networkID = spCookie.Value
	}

	http.SetCookie(w, &http.Cookie{
		Name:    "sp",
		Value:   networkID,
		Expires: time.Now().AddDate(1, 0, 0),
		Domain:  os.Getenv("COOKIE_DOMAIN"),
	})

	// TODO set P3P header

	// TODO handle GET request

	switch request.Method {
	case "POST":
		c.servePost(w, request, networkID)
	default:
		w.WriteHeader(http.StatusForbidden)
		w.Write(nil) // TODO make nice message
	}
}

func (c *collector) servePost(
	w http.ResponseWriter,
	request *http.Request,
	networkID string,
) {
	bodyBytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil) // TODO make nice message
		return
	}

	// this deserialization is used to check if we have any real data. It’s
	// slightly wasteful, but it’s certainly not a deal breaker right now and
	// the savings we get from not shipping empty events is huge

	trackerPayload := snowplow.TrackerPayload{}
	if err := json.Unmarshal(bodyBytes, &trackerPayload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil) // TODO make nice message
		return
	}

	if len(trackerPayload.Data) > 0 {
		collectorPayload := snowplow.CollectorPayload{
			Schema:        snowplow.CollectorPayloadSchema,
			IPAddress:     realRemoteAddr(request),
			Timestamp:     time.Now().UnixNano() / 1000000,
			Collector:     "Snowblower/0.0.1",
			UserAgent:     request.UserAgent(),
			Body:          string(bodyBytes),
			Headers:       requestHeadersAsArray(request),
			NetworkUserID: networkID,
			//Encoding:      "",
			//RefererURI:    "",
			//Path:          "",
			//QueryString:   "",
			//ContentType:   "",
			//Hostname:      "",
		}
		messageBytes, err := json.Marshal(collectorPayload)
		if err != nil {
			log.Printf("Error marshalling JSON: %s", err)
			return
		}
		message := string(messageBytes)
		c.publisher.publish(message)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}

func startCollector() {

	healthHandler := &health{}
	http.Handle("/api/health", healthHandler)
	http.Handle("/health", healthHandler)

	collectorHandler := &collector{
		publisher: &SNSPublisher{
			service: config.snsService,
			topic:   config.snsTopic,
		},
	}

	http.Handle("/com.snowplowanalytics.snowplow/tp2", collectorHandler)
	http.Handle("/i", collectorHandler)

	portString := fmt.Sprintf(":%s", config.collectorPort)
	log.Printf("Starting server on %s", portString)
	http.ListenAndServe(portString, nil)
}
