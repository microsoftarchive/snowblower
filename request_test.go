package main

import (
	"net/http"
	"testing"
)

func makeRequest() *http.Request {
	request, _ := http.NewRequest("GET", "/", nil)
	request.RemoteAddr = "1.2.3.4"
	return request
}

func TestRealRemoteAddrWithNoXFFHeader(t *testing.T) {
	request := makeRequest()
	addr := realRemoteAddr(request)
	if addr != "1.2.3.4" {
		t.Errorf("Expected 1.2.3.4, got: %s", addr)
	}
}

func TestRealRemoteAddrWithSimpleXFFHeader(t *testing.T) {
	request := makeRequest()
	request.Header.Set("X-Forwarded-For", "9.9.9.9")
	addr := realRemoteAddr(request)
	if addr != "9.9.9.9" {
		t.Errorf("Expected 9.9.9.9, got: %s", addr)
	}
}

func TestRealRemoteAddrWithJunkXFFHeader(t *testing.T) {
	request := makeRequest()
	request.Header.Set("X-Forwarded-For", "JUNK")
	addr := realRemoteAddr(request)
	if addr != "1.2.3.4" {
		t.Errorf("Expected 1.2.3.4, got: %s", addr)
	}
}

func TestRealRemoteAddrWithComplexXFFHeader(t *testing.T) {
	request := makeRequest()
	request.Header.Set("X-Forwarded-For", "9.9.9.9, 8.8.8.8")
	addr := realRemoteAddr(request)
	if addr != "9.9.9.9" {
		t.Errorf("Expected 9.9.9.9, got: %s", addr)
	}
}

func TestRealRemoteAddrWithPrivate10DotXFFHeader(t *testing.T) {
	request := makeRequest()
	request.Header.Set("X-Forwarded-For", "10.0.0.2, 8.8.8.8")
	addr := realRemoteAddr(request)
	if addr != "8.8.8.8" {
		t.Errorf("Expected 8.8.8.8, got: %s", addr)
	}
}

func TestRealRemoteAddrWithPrivate172DotXFFHeader(t *testing.T) {
	request := makeRequest()
	request.Header.Set("X-Forwarded-For", "172.16.0.2, 8.8.8.8")
	addr := realRemoteAddr(request)
	if addr != "8.8.8.8" {
		t.Errorf("Expected 8.8.8.8, got: %s", addr)
	}
}

func TestRealRemoteAddrWithPrivate192DotXFFHeader(t *testing.T) {
	request := makeRequest()
	request.Header.Set("X-Forwarded-For", "192.168.2.2, 8.8.8.8")
	addr := realRemoteAddr(request)
	if addr != "8.8.8.8" {
		t.Errorf("Expected 8.8.8.8, got: %s", addr)
	}
}

func TestRequestHeadersAsArray(t *testing.T) {
	request := makeRequest()
	request.Header.Add("foo", "bar")
	request.Header.Add("foo", "baz")
	headers := requestHeadersAsArray(request)
	if len(headers) != 2 {
		t.Errorf("Wrong number of headers, expected 2, got %v", len(headers))
	}
}
