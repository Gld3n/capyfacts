package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	app := application{logger: logger}

	_, err = pgx.Connect(context.Background(), os.Getenv("DB_PATH"))
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	slog.Info("connection established successfully")
	slog.Info("starting server on port 8080")
	err = http.ListenAndServe("localhost:8080", app.routes())
	if err != nil {
		slog.Error(err.Error())
	}
}
