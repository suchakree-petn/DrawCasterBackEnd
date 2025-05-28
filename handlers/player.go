package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"GolangPractice/database"
	"GolangPractice/models"
)

func SaveGameData(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var gameData models.GameDataDB
	err := json.NewDecoder(r.Body).Decode(&gameData)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	err = database.InsertGameData(gameData)
	if err != nil {
		log.Println("DB Error:", err)
		http.Error(w, "Failed to save data", http.StatusInternalServerError)
		return
	}
	fmt.Fprintf(w, "Saved gamedata: %s  %d", gameData.PlayerId, gameData.Level)

}

func LoadGameData(w http.ResponseWriter, r *http.Request) {
	playerId := r.URL.Query().Get("playerId")
	if playerId == "" {
		http.Error(w, "Missing playerId", http.StatusBadRequest)
		return
	}

	data, err := database.GetGameData(playerId)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			// ไม่เจอ playerId => ส่ง 404 พร้อมข้อความ
			http.Error(w, "Player data not found", http.StatusNotFound)
			return
		}
		// error อื่น ๆ
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
