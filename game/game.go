package game

import (
	"tic-tac-toe/models"
)

func MakeMove(cg *models.Game, p *models.Player) bool {

	if p.Row < 0 || p.Row > 2 || p.Col < 0 || p.Col > 2 ||
		cg.Board[p.Row][p.Col] != "" {
		return false
	}

	switch p.CurrentPlayer {
	case cg.PlayerOneSymbol:
		cg.Board[p.Row][p.Col] = "X"
	case cg.PlayerTwoSymbol:
		cg.Board[p.Row][p.Col] = "O"
	default:
		return false
	}

	return true
}

func CheckWin(board [3][3]string) string {
	for row := 0; row < 3; row++ {
		if board[row][0] != "" &&
			board[row][0] == board[row][1] &&
			board[row][1] == board[row][2] {
			return board[row][0]
		}
	}

	for col := 0; col < 3; col++ {
		if board[0][col] != "" &&
			board[0][col] == board[1][col] &&
			board[1][col] == board[2][col] {
			return board[0][col]
		}
	}

	if board[0][0] != "" &&
		board[0][0] == board[1][1] &&
		board[1][1] == board[2][2] {
		return board[0][0]
	}

	if board[0][2] != "" &&
		board[0][2] == board[1][1] &&
		board[1][1] == board[2][0] {
		return board[0][2]
	}

	return ""
}

func isDraw(board [3][3]string) bool {
	for _, row := range board {
		for _, cell := range row {
			if cell == "" {
				return false
			}
		}
	}
	return true
}

func ClearBoard(board *[3][3]string) {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			board[i][j] = ""
		}
	}
}

func StatusGame(winner string, currentGame *models.Game) models.Game {
	switch {
	case winner != "":
		currentGame.Status = "finished"
		currentGame.Winner = winner
	case isDraw(currentGame.Board):
		currentGame.Status = "finished"
		currentGame.Winner = "draw"
	default:
		switch currentGame.СurrentPlayer {
		case "X":
			currentGame.СurrentPlayer = "O"
		default:
			currentGame.СurrentPlayer = "X"
		}
	}
	return *currentGame
}
