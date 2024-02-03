package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type Player struct {
	Name   string `json:"name"`
	Number int    `json:"number"`
}

var players = []Player{
	{Name: "Lionel Messi", Number: 10},
	{Name: "Gerard Piqué", Number: 3},
	{Name: "Ansu Fati", Number: 22},
}

func GetPlayers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(players)
}

func GetPlayer(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range players {
		if item.Name == params["name"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Player not found", http.StatusNotFound)
}

func GetHealth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("Barcelona App - Author: Your Name"))
}
