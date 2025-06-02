package models

type Player struct {
	ID            int    `json:"id"`
	CurrentPlayer string `json:"currentPlayer"`
	Row           int    `json:"row"`
	Col           int    `json:"col"`
}
