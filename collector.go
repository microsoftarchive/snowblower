package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"code.google.com/p/go-uuid/uuid"
)

type collector struct {
	publisher EventPublisher
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
	bytes, err := ioutil.ReadAll(request.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil) // TODO make nice message
		return
	}

	trackerPayload := TrackerPayload{}
	if err := json.Unmarshal(bytes, &trackerPayload); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(nil) // TODO make nice message
		return
	}

	for _, event := range trackerPayload.Data {
		if c.publisher != nil {
			event.CollectorTimestamp = string(time.Now().UnixNano())
			event.CollectorName = "Snowblower"
			event.CollectorVersion = "0.0.1"
			event.UserIPAddress = realRemoteAddr(request)
			event.NetworkUserID = networkID
			c.publisher.publish(&event)
		} else {
			log.Printf("Sending event to /dev/null: %s", event)
		}
	}
	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}
