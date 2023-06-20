package handlers

import (
	"github.com/go-chi/chi/v5"
	"github.com/henriquemdimer/go-crud/auth"
	controllers "github.com/henriquemdimer/go-crud/controllers"
)

func LoadUserRoutes(router chi.Router) {
	router.Post("/", controllers.CreateUser)
	router.Post("/login", controllers.Login)

	infos := router.Group(nil)
	infos.Use(auth.Middleware)
	infos.Get("/", controllers.GetUser)
}
