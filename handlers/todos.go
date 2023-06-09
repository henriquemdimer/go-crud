package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/henriquemdimer/go-crud/controllers"
)

func LoadTodoRoutes(router chi.Router) {
	router.Post("/", controllers.CreateTodo)
	router.Get("/", controllers.ReadAllTodos)
	router.Get("/{todoID}", controllers.ReadOneTodo)
	router.Delete("/{todoID}", controllers.DeleteTodo)
	router.Put("/{todoID}", controllers.UpdateTodo)
}
