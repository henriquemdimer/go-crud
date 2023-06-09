package todoControllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	todoModel "github.com/henriquemdimer/go-crud/models/todo"
)

func Update(w http.ResponseWriter, r *http.Request) {
	todoID := chi.URLParam(r, "todoID")
	id, _ := strconv.ParseInt(todoID, 0, 0)

	var todo todoModel.Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	rows, err := todoModel.UpdateOne(id, todo)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fmt.Sprintf("Todos updated: %d", rows))
}
