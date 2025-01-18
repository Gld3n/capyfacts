package main

import (
	"fmt"
	"net/http"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CapyFacts' on air!")
}

func (app *application) getAllFactsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "All CapyFacts")
}

func (app *application) getRandomFactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "A random CapyFact")
}

func (app *application) createFactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CapyFact created")
}

func (app *application) updateFactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CapyFact updated")
}

func (app *application) deleteFactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CapyFact deleted")
}
