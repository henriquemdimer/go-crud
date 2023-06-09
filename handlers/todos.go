package handlers

import (
	"github.com/go-chi/chi/v5"
	todoControllers "github.com/henriquemdimer/go-crud/controllers/todo"
)

func LoadTodoRoutes(router chi.Router) {
	router.Post("/", todoControllers.Create)
	router.Get("/", todoControllers.ReadAll)
	router.Get("/{todoID}", todoControllers.ReadOne)
	router.Delete("/{todoID}", todoControllers.Delete)
	router.Put("/{todoID}", todoControllers.Update)
}
