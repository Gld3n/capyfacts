package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func serveJSONResponse(w http.ResponseWriter, val any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(val); err != nil {
		return fmt.Errorf("error encoding json: %s", err.Error())
	}

	return nil
}

func errorResponse(w http.ResponseWriter, err error, statusCode int) {
	http.Error(w, err.Error(), statusCode)
}

func newMessage(m string) *map[string]string {
	msg := make(map[string]string, 1)
	msg["message"] = m
	return &msg
}
