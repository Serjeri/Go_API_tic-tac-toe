package game

import "tic-tac-toe/models"

func MakeMove(board *[3][3]string, p *models.Player) bool {

	if p.Row < 0 || p.Row > 2 || p.Col < 0 || p.Col > 2 ||
		board[p.Row][p.Col] != " " {
		return false
	}

	board[p.Row][p.Col] = p.CurrentPlayer

	if checkWin(*board, p.CurrentPlayer, p.Row, p.Col) {
		return true
	}

	if isDraw(*board) {
		return true
	}

	if p.CurrentPlayer == "X" {
		p.CurrentPlayer = "O"
	} else {
		p.CurrentPlayer = "X"
	}

	return true
}

func checkWin(board [3][3]string, currentPlayer string, row, col int) bool {

	if board[row][0] == currentPlayer &&
		board[row][1] == currentPlayer &&
		board[row][2] == currentPlayer {
		return true
	}

	if board[0][col] == currentPlayer &&
		board[1][col] == currentPlayer &&
		board[2][col] == currentPlayer {
		return true
	}

	if (row == col && board[0][0] == currentPlayer &&
		board[1][1] == currentPlayer &&
		board[2][2] == currentPlayer) ||
		(row+col == 2 && board[0][2] == currentPlayer &&
			board[1][1] == currentPlayer &&
			board[2][0] == currentPlayer) {
		return true
	}

	return false
}

func isDraw(board [3][3]string) bool {
	for _, row := range board {
		for _, cell := range row {
			if cell == " " {
				return false
			}
		}
	}
	return true
}

func ClearBoard(board *[3][3]string) {
	for i := range board {
		for j := range board[i] {
			board[i][j] = " "
		}
	}
}
