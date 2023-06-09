package handlers

import (
	"github.com/go-chi/chi/v5"
	userControllers "github.com/henriquemdimer/go-crud/controllers/user"
)

func LoadUserRoutes(router chi.Router) {
	router.Post("/", userControllers.Create)
}
