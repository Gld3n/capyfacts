package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Capyfacts on air!")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /", http.HandlerFunc(homeHandler))
	slog.Info("starting server on port 8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		slog.Error(err.Error())
	}
}
