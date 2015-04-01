package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

const emptyPost = `
{"schema":"iglu:com.snowplowanalytics.snowplow/payload_data/jsonschema/1-0-0","data":[]}
`

const singlePost = `
{"schema":"iglu:com.snowplowanalytics.snowplow/payload_data/jsonschema/1-0-0","data":[{"aid":"Wunderlist/3.2.2 fd67664087d7168d5f8538186441577e241ecb78","res":"750x1334","uid":"0","p":"mob","cx":"eyJzY2hlbWEiOiJpZ2x1OmNvbS5zbm93cGxvd2FuYWx5dGljcy5zbm93cGxvd1wvY29udGV4dHNcL2pzb25zY2hlbWFcLzEtMC0wIiwiZGF0YSI6W3sic2NoZW1hIjoiaWdsdTpjb20uc25vd3Bsb3dhbmFseXRpY3Muc25vd3Bsb3dcL21vYmlsZV9jb250ZXh0XC9qc29uc2NoZW1hXC8xLTAtMSIsImRhdGEiOnsib3NUeXBlIjoiaW9zIiwibmV0d29ya1R5cGUiOiJtb2JpbGUiLCJvcGVuSWRmYSI6IkJBMURBOTI3LTQ1MkEtNjZEMy0zREY1LUFFMDJCMTQ2OEMwQiIsIm9zVmVyc2lvbiI6IjguMiIsIm5ldHdvcmtUZWNobm9sb2d5IjoiQ1RSYWRpb0FjY2Vzc1RlY2hub2xvZ3lMVEUiLCJhcHBsZUlkZnYiOiIxRjBFNDUzNC1FODUyLTQ5OTMtOTI3Mi01MEM2NjEzOEYxRUQiLCJjYXJyaWVyIjoiQVQmVCIsImRldmljZU1hbnVmYWN0dXJlciI6IkFwcGxlIEluYy4iLCJhcHBsZUlkZmEiOiIwMDJDMTM0MS1FOUQwLTQ2NTEtOTY2OS00NkZEOTNEMjgzRkQiLCJkZXZpY2VNb2RlbCI6ImlQaG9uZSJ9fV19","dtm":"1427898727829","tv":"ios-0.3.2","tna":"com.wunderlist","ue_px":"eyJzY2hlbWEiOiJpZ2x1OmNvbS5zbm93cGxvd2FuYWx5dGljcy5zbm93cGxvd1wvdW5zdHJ1Y3RfZXZlbnRcL2pzb25zY2hlbWFcLzEtMC0wIiwiZGF0YSI6eyJzY2hlbWEiOiJpZ2x1OmNvbS53dW5kZXJsaXN0XC9jbGllbnRfZXZlbnRcL2pzb25zY2hlbWFcLzEtMC0wIiwiZGF0YSI6eyJldmVudCI6ImNsaWVudC5hY2NvdW50LmxvZ2luIiwicGFyYW1ldGVycyI6eyJ0eXBlIjoicGFzc3dvcmQifX19fQ","e":"ue","lang":"en","vp":"750x1334","eid":"8dbe2937-beea-45ee-9e78-269343ba6561"}]}
`

func TestHealth(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/health", nil)
	recorder := httptest.NewRecorder()
	healthHandler(recorder, request)
	if code := recorder.Code; code != http.StatusOK {
		t.Errorf("expected call to be successul. Got %d instead", code)
	}
	if contentType := recorder.Header().Get("Content-Type"); contentType != "application/json" {
		t.Errorf("Expected Content-Type to equal application/json. Got %s instead", contentType)
	}
}

func TestTrackingPostRejectsGet(t *testing.T) {
	request, _ := http.NewRequest("GET", trackingPostPath, nil)
	recorder := httptest.NewRecorder()
	trackingPostHandler(recorder, request)
	if code := recorder.Code; code != http.StatusForbidden {
		t.Errorf("expected call to be forbidden. Got %d instead", code)
	}
}

func trackingPost(body string) (*httptest.ResponseRecorder, *http.Request) {
	request, _ := http.NewRequest(
		"POST", trackingPostPath, strings.NewReader(body))
	recorder := httptest.NewRecorder()
	trackingPostHandler(recorder, request)
	return recorder, request
}

func TestEmptyTrackingPostReturnsOK(t *testing.T) {
	recorder, _ := trackingPost(emptyPost)
	if code := recorder.Code; code != http.StatusOK {
		t.Errorf("expected call to be successul. Got %d instead", code)
	}
}

// func TestSingleTrackingPostReturnsOK(t *testing.T) {
// 	setup()
// 	recorder, _ := trackingPost(singlePost)
// 	if code := recorder.Code; code != http.StatusOK {
// 		t.Errorf("expected call to be successul. Got %d instead", code)
// 	}
// }
