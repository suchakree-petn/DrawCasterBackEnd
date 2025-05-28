package models

type GameDataDB struct {
	PlayerId string `json:"playerid"`
	Level    int    `json:"level"`
	Gold     int    `json:"gold"`
}
