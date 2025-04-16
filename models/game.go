package models

type Game struct {
	ID              int          `json:"id"`
	Board           [3][3]string `json:"board"`
	Status          string       `json:"status"`
	Winner          string       `json:"winner"`
	Draw            string       `json:"draw"`
	СurrentPlayer   string       `json:"сurrentPlayer"`
	PlayerOneSymbol string       `json:"playerOneSymbol"`
	PlayerTwoSymbol string       `json:"playerTwoSymbol"`
}
