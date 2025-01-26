package main

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/gld3n/capyfacts/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
)

type application struct {
	logger *slog.Logger
	facts  models.FactsModelInterface
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	err := godotenv.Load()
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	conn, err := pgxpool.New(context.Background(), os.Getenv("DB_PATH"))
	if err != nil {
		logger.Error(fmt.Sprintf("Unable to establish database connection: %s", err.Error()))
		os.Exit(1)
	}
	defer conn.Close()

	app := &application{logger: logger, facts: &models.FactsModel{DB: conn}}

	slog.Info("connection established successfully")
	slog.Info("starting server on port 8080")
	err = http.ListenAndServe("localhost:8080", app.routes())
	if err != nil {
		slog.Error(err.Error())
	}
}
