package main

import (
	"fmt"
	"github.com/gld3n/capyfacts/internal/models"
	"net/http"
)

func (app *application) homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CapyFacts' on air!")
}

func (app *application) getAllFactsHandler(w http.ResponseWriter, r *http.Request) {
	facts, err := app.facts.GetAll(models.Behavior, 10, 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = serveJSONResponse(w, facts); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) getRandomFactHandler(w http.ResponseWriter, r *http.Request) {
	fact, err := app.facts.Random()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = serveJSONResponse(w, fact); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
