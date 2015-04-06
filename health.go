package main

import (
	"encoding/json"
	"net/http"
)

type health struct{}

func (h *health) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	msg := make(map[string]string)
	msg["up"] = "true"
	output, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(output)
}
