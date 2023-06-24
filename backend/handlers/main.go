package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/henriquemdimer/go-crud/controllers"
)

func LoadMainRoutes(router chi.Router) {
	router.Get("/", controllers.PingMain)
}
