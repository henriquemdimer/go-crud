package handlers

import (
	"github.com/go-chi/chi/v5"
	controllers "github.com/henriquemdimer/go-crud/controllers"
)

func LoadUserRoutes(router chi.Router) {
	router.Post("/", controllers.CreateUser)
}
