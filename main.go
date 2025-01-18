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
	slog.Info("connection established successfully")

	slog.Info("starting server on port 8080")
	err := http.ListenAndServe("localhost:8080", routes())
	if err != nil {
		slog.Error(err.Error())
	}
}
