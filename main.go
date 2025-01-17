package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CapyFacts' on air!")
}

func getAllFactsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All CapyFacts")
}

func getRandomFactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A random CapyFact")
}

func createFactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CapyFact created")
}

func updateFactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CapyFact updated")
}

func deleteFactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CapyFact deleted")
}

func main() {
	mux := http.NewServeMux()
	mux.Handle("GET /", http.HandlerFunc(homeHandler))
	mux.Handle("GET /facts", http.HandlerFunc(getAllFactsHandler))
	mux.Handle("GET /facts/random", http.HandlerFunc(getRandomFactHandler))
	mux.Handle("POST /facts", http.HandlerFunc(createFactHandler))
	mux.Handle("PUT /facts{id}", http.HandlerFunc(updateFactHandler))
	mux.Handle("DELETE /facts{id}", http.HandlerFunc(deleteFactHandler))

	slog.Info("starting server on port 8080")
	err := http.ListenAndServe("localhost:8080", mux)
	if err != nil {
		slog.Error(err.Error())
	}
}
