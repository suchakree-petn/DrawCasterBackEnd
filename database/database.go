package database

import (
	"GolangPractice/models"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	}

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}
	err = DB.Ping()
	if err != nil {
		log.Fatalf("Cannot reach the database: %v", err)
	}
	log.Println("Database connection established successfully")
}

func InsertGameData(data models.GameDataDB) error {
	query := `
		INSERT INTO "GameData" (playerid, level, gold)
		VALUES ($1, $2, $3)
		ON CONFLICT (playerid)
		DO UPDATE SET level = EXCLUDED.level, gold = EXCLUDED.gold;
	`
	_, err := DB.Exec(query, data.PlayerId, data.Level, data.Gold)
	return err
}

func GetGameData(playerId string) (models.GameDataDB, error) {
	var data models.GameDataDB
	query := `SELECT playerid, level, gold FROM "GameData" WHERE playerid = $1`
	err := DB.QueryRow(query, playerId).Scan(&data.PlayerId, &data.Level, &data.Gold)

	if err != nil {
		if err == sql.ErrNoRows {
			// ไม่พบข้อมูลผู้เล่น
			return data, fmt.Errorf("playerId '%s' not found", playerId)
		}
		// error อื่น ๆ เช่น database ล่ม
		return data, err
	}

	return data, nil
}
