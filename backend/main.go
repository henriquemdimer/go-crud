package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/henriquemdimer/go-crud/db"
	"github.com/henriquemdimer/go-crud/handlers"
	"github.com/rs/cors"
)

func main() {
	db, err := db.Open()
	if err != nil {
		fmt.Println("Error connecting db:", err)
		os.Exit(2)
	}
	db.Close()

	router := chi.NewRouter()

	router.Route("/todos", handlers.LoadTodoRoutes)
	router.Route("/users", handlers.LoadUserRoutes)

	handler := cors.Default().Handler(router)

	http.ListenAndServe(":8080", handler)
}
