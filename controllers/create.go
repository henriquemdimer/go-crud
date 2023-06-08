package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	model_todo "github.com/henriquemdimer/go-crud/models/todo"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var todo model_todo.Todo

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		log.Println("Error parsing payload:", err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	id, err := model_todo.InsertOne(todo)
	var resp map[string]any

	if err != nil {
		resp = map[string]any{
			"Error":   true,
			"Message": fmt.Sprintf("Error ocurred while inserting todo: %v", err),
		}
	} else {
		resp = map[string]any{
			"Error":   false,
			"Message": fmt.Sprintf("Todo inserted with ID: %d", id),
		}
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
