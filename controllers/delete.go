package controllers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	todoModel "github.com/henriquemdimer/go-crud/models/todo"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	id, _ := strconv.ParseInt(todoID, 0, 0)

	rows, err := todoModel.DeleteOne(id)
	if err != nil || rows > 1 {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(rows)
}
