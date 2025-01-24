package main

import (
	"encoding/json"
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
	var fact *models.Fact

	if err := json.NewDecoder(r.Body).Decode(&fact); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err := app.facts.Create(fact)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp := make(map[string]any)
	resp["fact_created"] = fact

	if err = serveJSONResponse(w, resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (app *application) updateFactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CapyFact updated")
}

func (app *application) deleteFactHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "CapyFact deleted")
}
