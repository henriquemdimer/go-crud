package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/henriquemdimer/go-crud/models"
)

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	userID := r.Context().Value("userID").(int64)

	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		returnError(w, err)
		return
	}

	id, err := models.InsertOneTodo(todo, userID)
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

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	id, _ := strconv.ParseInt(todoID, 0, 0)
	userID := r.Context().Value("userID").(int64)

	rows, err := models.DeleteOneTodo(id, userID)
	if err != nil || rows > 1 {
		returnError(w, err)
		return
	}

	if rows < 1 {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmt.Sprintf("Todos deleted: %d", rows))
}

func ReadAllTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	userID := r.Context().Value("userID").(int64)
	todos, err := models.GetAllTodos(userID)

	if err != nil {
		returnError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todos)
}

func ReadOneTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	id, _ := strconv.ParseInt(todoID, 0, 0)
	userID := r.Context().Value("userID").(int64)

	todo, err := models.GetOneTodo(id, userID)
	if err != nil {
		returnError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	id, _ := strconv.ParseInt(todoID, 0, 0)
	userID := r.Context().Value("userID").(int64)

	var todo models.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		returnError(w, err)
		return
	}

	rows, err := models.UpdateOneTodo(id, todo.Done, userID)
	if err != nil {
		returnError(w, err)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmt.Sprintf("Todos updated: %d", rows))
}
