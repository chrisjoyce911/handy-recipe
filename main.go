package main

import (
	"fmt"
	"log"
	"net/http"

	"handy-recipe/handlers"
	"handy-recipe/store"
)

func main() {
	// Load the JSONStore from the directory containing individual recipe JSON files.
	jsonStore, err := store.NewJSONStoreFromDir("data")
	if err != nil {
		log.Fatalf("Error loading recipes: %v", err)
	}

	// Pass the store to your handlers.
	h := handlers.NewHandler(jsonStore)
	mux := h.SetupRoutes()

	fmt.Println("🚀 Server is running on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
