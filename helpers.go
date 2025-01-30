package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
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

func parseAndValidateLimit(queryLimit string, targetLimit *int) error {
	const minLimit = 1
	const maxLimit = 100

	limit, err := strconv.Atoi(queryLimit)
	if err != nil {
		return fmt.Errorf("invalid limit provided: '%d'", targetLimit)
	}
	if limit < minLimit || limit > maxLimit {
		return errors.New("maximum allowed limit is 100")
	}

	*targetLimit = limit
	return nil
}

func parseAndValidateOffset(queryOffset string, targetOffset *int) error {
	ofs, err := strconv.Atoi(queryOffset)
	if err != nil {
		return fmt.Errorf("invalid offset provided: '%d'", targetOffset)
	}

	*targetOffset = ofs
	return nil
}
