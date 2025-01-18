package main

import (
	"net/http"
)

func (app *application) routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /", http.HandlerFunc(app.homeHandler))
	mux.Handle("GET /facts", http.HandlerFunc(app.getAllFactsHandler))
	mux.Handle("GET /facts/random", http.HandlerFunc(app.getRandomFactHandler))
	mux.Handle("POST /facts", http.HandlerFunc(app.createFactHandler))
	mux.Handle("PUT /facts/{id}", http.HandlerFunc(app.updateFactHandler))
	mux.Handle("DELETE /facts/{id}", http.HandlerFunc(app.deleteFactHandler))

	return mux
}
