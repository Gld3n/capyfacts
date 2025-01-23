package main

import (
	"encoding/json"
	"net/http"
)

func serveJSONResponse(w http.ResponseWriter, val any) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(val)
}
