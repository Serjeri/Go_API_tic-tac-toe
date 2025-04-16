package models

type Message struct {
	ID         int          `json:"id"`
	Board      [3][3]string `json:"board"`
	NextPlayer string       `json:"nextPlayer"`
	GameStatus string       `json:"gameStatus"`
	Winner     string       `json:"winner"`
	Message    string       `json:"message"`
}
