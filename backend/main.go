package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/henriquemdimer/go-crud/db"
	"github.com/henriquemdimer/go-crud/handlers"
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

	http.ListenAndServe(":8080", router)
}
