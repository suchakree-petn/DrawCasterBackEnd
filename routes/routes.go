package routes

import (
	"GolangPractice/handlers"
	"net/http"
)

func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/saveGameData", handlers.SaveGameData)
	mux.HandleFunc("/loadGameData", handlers.LoadGameData)
}
