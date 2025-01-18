package main

import "net/http"

func routes() *http.ServeMux {
	mux := http.NewServeMux()
	mux.Handle("GET /", http.HandlerFunc(homeHandler))
	mux.Handle("GET /facts", http.HandlerFunc(getAllFactsHandler))
	mux.Handle("GET /facts/random", http.HandlerFunc(getRandomFactHandler))
	mux.Handle("POST /facts", http.HandlerFunc(createFactHandler))
	mux.Handle("PUT /facts/{id}", http.HandlerFunc(updateFactHandler))
	mux.Handle("DELETE /facts/{id}", http.HandlerFunc(deleteFactHandler))

	return mux
}
