package main

import (
	"GolangPractice/database"
	"GolangPractice/routes"
	"log"
	"net/http"
	"os"
)

func main() {
	// Init database and routes
	database.InitDB()
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	// Default test endpoint
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Backend Running"))
	})

	// Get PORT from Render (or fallback)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	log.Println("Server is running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, mux))
}
