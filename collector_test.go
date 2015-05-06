package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"encoding/json"
	"reflect"
	"sort"
)

const noDataPayload = `
{"schema":"iglu:com.snowplowanalytics.snowplow/payload_data/jsonschema/1-0-0","data":[]}
`

const singleDataPayload = `
{"schema":"iglu:com.snowplowanalytics.snowplow/payload_data/jsonschema/1-0-0","data":[{"aid":"Wunderlist/3.2.2 fd67664087d7168d5f8538186441577e241ecb78","res":"750x1334","uid":"0","p":"mob","cx":"eyJzY2hlbWEiOiJpZ2x1OmNvbS5zbm93cGxvd2FuYWx5dGljcy5zbm93cGxvd1wvY29udGV4dHNcL2pzb25zY2hlbWFcLzEtMC0wIiwiZGF0YSI6W3sic2NoZW1hIjoiaWdsdTpjb20uc25vd3Bsb3dhbmFseXRpY3Muc25vd3Bsb3dcL21vYmlsZV9jb250ZXh0XC9qc29uc2NoZW1hXC8xLTAtMSIsImRhdGEiOnsib3NUeXBlIjoiaW9zIiwibmV0d29ya1R5cGUiOiJtb2JpbGUiLCJvcGVuSWRmYSI6IkJBMURBOTI3LTQ1MkEtNjZEMy0zREY1LUFFMDJCMTQ2OEMwQiIsIm9zVmVyc2lvbiI6IjguMiIsIm5ldHdvcmtUZWNobm9sb2d5IjoiQ1RSYWRpb0FjY2Vzc1RlY2hub2xvZ3lMVEUiLCJhcHBsZUlkZnYiOiIxRjBFNDUzNC1FODUyLTQ5OTMtOTI3Mi01MEM2NjEzOEYxRUQiLCJjYXJyaWVyIjoiQVQmVCIsImRldmljZU1hbnVmYWN0dXJlciI6IkFwcGxlIEluYy4iLCJhcHBsZUlkZmEiOiIwMDJDMTM0MS1FOUQwLTQ2NTEtOTY2OS00NkZEOTNEMjgzRkQiLCJkZXZpY2VNb2RlbCI6ImlQaG9uZSJ9fV19","dtm":"1427898727829","tv":"ios-0.3.2","tna":"com.wunderlist","ue_px":"eyJzY2hlbWEiOiJpZ2x1OmNvbS5zbm93cGxvd2FuYWx5dGljcy5zbm93cGxvd1wvdW5zdHJ1Y3RfZXZlbnRcL2pzb25zY2hlbWFcLzEtMC0wIiwiZGF0YSI6eyJzY2hlbWEiOiJpZ2x1OmNvbS53dW5kZXJsaXN0XC9jbGllbnRfZXZlbnRcL2pzb25zY2hlbWFcLzEtMC0wIiwiZGF0YSI6eyJldmVudCI6ImNsaWVudC5hY2NvdW50LmxvZ2luIiwicGFyYW1ldGVycyI6eyJ0eXBlIjoicGFzc3dvcmQifX19fQ","e":"ue","lang":"en","vp":"750x1334","eid":"8dbe2937-beea-45ee-9e78-269343ba6561"}]}
`

type MockPublisher interface {
	messages() []string
}

type testPublisher struct {
	collectedMessages []string
}

func (p *testPublisher) publish(message string) {
	p.collectedMessages = append(p.collectedMessages, message)
}

func (p *testPublisher) messages() []string {
	return p.collectedMessages
}

func setupRequest(method string, path string, body string) (*httptest.ResponseRecorder, *http.Request) {
	request, _ := http.NewRequest(method, path, strings.NewReader(body))
	return httptest.NewRecorder(), request
}

func performTestRequest(
	recorder *httptest.ResponseRecorder,
	request *http.Request,
) *collector {
	p := &testPublisher{}
	c := &collector{publisher: p}
	c.ServeHTTP(recorder, request)
	return c
}

func performTestRequestAndGetPublisher(
	recorder *httptest.ResponseRecorder,
	request *http.Request,
	t *testing.T,
) MockPublisher {
	collector := performTestRequest(recorder, request)
	if p, ok := collector.publisher.(MockPublisher); ok {
		return p
	}
	t.Error("Internal problem with test interfaces")
	return nil
}

func unmarshalCollectorPayload(p MockPublisher, t *testing.T) CollectorPayload {
	message := p.messages()[0]
	messageBytes := []byte(message)
	collectorPayload := CollectorPayload{}

	if err := json.Unmarshal(messageBytes, &collectorPayload); err != nil {
		t.Error("Error unmarshalling collectorPayload after creation")
	}

	return collectorPayload
}

func TestSetsCookieWhenNoneSent(t *testing.T) {
	recorder, request := setupRequest("GET", "/", "")
	_ = performTestRequest(recorder, request)

	cookieValue := recorder.Header().Get("Set-Cookie")
	if cookieValue == "" {
		t.Error("Got no cookie")
	}
}

func TestSetsSameCookieAsSent(t *testing.T) {
	recorder, request := setupRequest("GET", "/", "")
	request.AddCookie(&http.Cookie{Name: "sp", Value: "TEST"})
	_ = performTestRequest(recorder, request)

	cookieValue := recorder.Header().Get("Set-Cookie")
	if !strings.HasPrefix(cookieValue, "sp=TEST") {
		t.Error("Didn’t get proper cookie back")
	}
}

func TestEmptyTrackingPostSendsNoMessages(t *testing.T) {
	recorder, request := setupRequest("POST", "/", "")
	publisher := performTestRequestAndGetPublisher(recorder, request, t)
	if len(publisher.messages()) == 1 {
		t.Errorf("Shit, got %v events", len(publisher.messages()))
	}
}

func TestNoDataTrackingPostSendsNoMessages(t *testing.T) {
	recorder, request := setupRequest("POST", "/", noDataPayload)
	publisher := performTestRequestAndGetPublisher(recorder, request, t)
	if len(publisher.messages()) == 1 {
		t.Errorf("Shit, got %v events", len(publisher.messages()))
	}
}

func TestSingleDataTrackingPostSendsOneEvent(t *testing.T) {
	recorder, request := setupRequest("POST", "/", singleDataPayload)
	publisher := performTestRequestAndGetPublisher(recorder, request, t)
	if len(publisher.messages()) != 1 {
		t.Errorf("Shit, got %v events", len(publisher.messages()))
	}
}

func TestBodyIsPassedCorrectly(t *testing.T) {
	recorder, request := setupRequest("POST", "/", singleDataPayload)
	publisher := performTestRequestAndGetPublisher(recorder, request, t)
	collectorPayload := unmarshalCollectorPayload(publisher, t)
	if collectorPayload.Body != singleDataPayload {
		t.Error("TrackerPayload is corrupted")
	}
}

func TestHeadersArePassedCorrectly(t *testing.T) {
	recorder, request := setupRequest("POST", "/", singleDataPayload)
	request.Header.Add("foo", "bar")
	request.Header.Add("baz", "qux")
	publisher := performTestRequestAndGetPublisher(recorder, request, t)
	collectorPayload := unmarshalCollectorPayload(publisher, t)

	if len(collectorPayload.Headers) != 2 {
		t.Errorf("Expecting 2 headers but received %v", len(collectorPayload.Headers))
	}

	expectedHeaders := requestHeadersAsArray(request)

	sort.Strings(collectorPayload.Headers)
	sort.Strings(expectedHeaders)

	if !reflect.DeepEqual(collectorPayload.Headers, expectedHeaders) {
		t.Error("Headers are not passed correctly")
	}
}
