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

var (
	GamesStorage = make(map[int]models.Game)
	gameID       = 0
)

func StartNewGame(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var symbol models.Game

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "errors", http.StatusUnauthorized)
		return
	}

	defer r.Body.Close()

	if err := json.Unmarshal(body, &symbol); err != nil {
		http.Error(w, "incorrect stroke format", http.StatusBadRequest)
		return
	}

	gameID++
	newBoard := board
	game.ClearBoard(&newBoard)

	newGame := models.Game{
		ID:              gameID,
		Board:           newBoard,
		Status:          "in_progress",
		СurrentPlayer:   "X",
		PlayerOneSymbol: symbol.PlayerOneSymbol,
		PlayerTwoSymbol: symbol.PlayerTwoSymbol,
	}

	GamesStorage[gameID] = newGame

	jsonResponse, err := json.Marshal(newGame)
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

	currentGame, exists := GamesStorage[player.ID]
	if !exists {
		http.Error(w, "Game not found", http.StatusNotFound)
		return
	}

	if currentGame.Status != "in_progress" {
		http.Error(w, "Game is already finished", http.StatusBadRequest)
		return
	}

	if !game.MakeMove(&currentGame, &player) {
		http.Error(w, "Is the cage already boarded up or have you gone off the board", http.StatusBadRequest)
		return
	}

	winner := game.CheckWin(currentGame.Board)
	statusGame := game.StatusGame(winner, &currentGame)

	GamesStorage[player.ID] = statusGame

	response := models.Message{
		ID:         player.ID,
		Board:      currentGame.Board,
		NextPlayer: currentGame.СurrentPlayer,
		GameStatus: currentGame.Status,
		Winner:     currentGame.Winner,
		Message:    "Move successful",
	}

	jsonResponse, err := json.Marshal(response)
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}
