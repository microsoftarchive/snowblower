package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHealth(t *testing.T) {
	request, _ := http.NewRequest("GET", "/api/health", nil)
	recorder := httptest.NewRecorder()

	h := &health{}
	h.ServeHTTP(recorder, request)

	if code := recorder.Code; code != http.StatusOK {
		t.Errorf("expected call to be successul. Got %d instead", code)
	}

	contentType := recorder.Header().Get("Content-Type")
	if contentType != "application/json" {
		t.Errorf("Expected application/json Content-Type. Got %s", contentType)
	}
}
