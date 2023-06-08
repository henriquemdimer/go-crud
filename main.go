package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/henriquemdimer/go-crud/controllers"
	"github.com/henriquemdimer/go-crud/db"
)

func main() {
	db, err := db.Open()
	if err != nil {
		fmt.Println("Error connecting db:", err)
		os.Exit(2)
	}
	db.Close()

	r := chi.NewRouter()

	r.Post("/", controllers.Create)
	r.Get("/", controllers.ReadAll)
	r.Get("/{todoID}", controllers.ReadOne)
	r.Delete("/{todoID}", controllers.Delete)

	http.ListenAndServe(":8080", r)
}
