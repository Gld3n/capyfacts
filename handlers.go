package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gld3n/capyfacts/internal/models"
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

type factRequest struct {
	Title    string `json:"title"`
	Content  string `json:"content"`
	Category string `json:"category"`
}

func (app *application) createFactHandler(w http.ResponseWriter, r *http.Request) {
	var factReq *factRequest

	if err := json.NewDecoder(r.Body).Decode(&factReq); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fact, err := models.NewFact(factReq.Title, factReq.Content, factReq.Category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = app.facts.Create(fact)
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
	idRequest := r.PathValue("id")

	id, err := strconv.Atoi(idRequest)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err = app.facts.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	msg := make(map[string]string, 1)
	msg["message"] = fmt.Sprintf("successfully deleted fact with id %d", id)

	if err = serveJSONResponse(w, msg); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
