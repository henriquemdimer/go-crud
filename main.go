package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/henriquemdimer/go-crud/controllers"
)

func main() {
	r := chi.NewRouter()

	r.Post("/", controllers.Create)
	r.Get("/", controllers.ReadAll)
	r.Get("/{todoID}", controllers.ReadOne)

	http.ListenAndServe(":8080", r)
}
