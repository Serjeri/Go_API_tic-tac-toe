package handlers

import (
	"encoding/json"
	"io"
	"net/http"
	"tic-tac-toe/game"
	"tic-tac-toe/models"
)

var board = [3][3]string{
	{" ", " ", " "},
	{" ", " ", " "},
	{" ", " ", " "},
}

func StartGame(w http.ResponseWriter, r *http.Request) {
	response := models.Player{
		CurrentPlayer: "X",
	}

	game.ClearBoard(&board)

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Failed to encode JSON", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func Move(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var player models.Player

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "errors", http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()

	if err := json.Unmarshal(body, &player); err != nil {
		http.Error(w, "incorrect stroke format", http.StatusBadRequest)
		return
	}
	status := game.MakeMove(&board, &player)
	if !status {
		http.Error(w, "Is the cage already boarded up or have you gone off the board", http.StatusBadRequest)
		return
	}

	response := models.Message{
		NextPlayer: player.CurrentPlayer,
		Message:    "Move successful",
	}

	jsonResponse, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
