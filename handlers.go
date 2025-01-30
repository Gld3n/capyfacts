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
	var c models.Category
	limit := 10
	offset := 0

	queryValues := r.URL.Query()
	limitParam := queryValues.Get("limit")
	offsetParam := queryValues.Get("offset")
	categoryParam := queryValues.Get("category")

	if limitParam != "" {
		err := parseAndValidateLimit(limitParam, &limit)
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}
	}

	if offsetParam != "" {
		err := parseAndValidateOffset(offsetParam, &offset)
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}
	}

	if len(categoryParam) > 0 {
		cat, err := models.ValidateCategory(categoryParam)
		if err != nil {
			errorResponse(w, err, http.StatusBadRequest)
			return
		}
		c = *cat
	}

	facts, err := app.facts.GetAll(c, limit, offset)
	if err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if err = serveJSONResponse(w, facts); err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
	}
}

func (app *application) getRandomFactHandler(w http.ResponseWriter, r *http.Request) {
	fact, err := app.facts.Random()
	if err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	if err = serveJSONResponse(w, fact); err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
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
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	if len(factReq.Title) < 10 {
		errorResponse(w, fmt.Errorf("title must be at least 16 characters"), http.StatusBadRequest)
		return
	}
	if len(factReq.Content) < 64 {
		errorResponse(w, fmt.Errorf("content must be at least 64 characters"), http.StatusBadRequest)
		return
	}

	fact, err := models.NewFact(factReq.Title, factReq.Content, factReq.Category)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	err = app.facts.Create(fact)
	if err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	resp := make(map[string]any)
	resp["fact_created"] = fact

	if err = serveJSONResponse(w, resp); err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
	}
}

func (app *application) updateFactHandler(w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")

	id, err := strconv.Atoi(idRequest)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	var factReq *factRequest

	if err = json.NewDecoder(r.Body).Decode(&factReq); err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	fact, err := models.NewFact(factReq.Title, factReq.Content, factReq.Category)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}
	fact.ID = id

	if err = app.facts.Edit(fact); err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	msg := newMessage(fmt.Sprintf("successfully updated fact with id %d", id))

	if err = serveJSONResponse(w, msg); err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
	}
}

func (app *application) deleteFactHandler(w http.ResponseWriter, r *http.Request) {
	idRequest := r.PathValue("id")

	id, err := strconv.Atoi(idRequest)
	if err != nil {
		errorResponse(w, err, http.StatusBadRequest)
		return
	}

	if err = app.facts.Delete(id); err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
		return
	}

	msg := newMessage(fmt.Sprintf("successfully deleted fact with id %d", id))

	if err = serveJSONResponse(w, msg); err != nil {
		errorResponse(w, err, http.StatusInternalServerError)
	}
}
