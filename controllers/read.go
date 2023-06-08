package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/henriquemdimer/go-crud/models"
)

func ReadAll(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	todos, err := models.GetAll()
	if err != nil {
		log.Println("Error while reading todos:", err)
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func ReadOne(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	id, _ := strconv.ParseInt(todoID, 0, 0)
	todo, err := models.GetOne(id)
	if err != nil {
		log.Println("Error:", err)
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}
