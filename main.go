package main

import (
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
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
	err := godotenv.Load()
	if err != nil {
		slog.Error(err.Error())
	}

	_, err = pgx.Connect(context.Background(), os.Getenv("DB_PATH"))
	if err != nil {
		slog.Error(err.Error())
	}
	slog.Info("connection established successfully")

	slog.Info("starting server on port 8080")
	err = http.ListenAndServe("localhost:8080", routes())
	if err != nil {
		slog.Error(err.Error())
	}
}
