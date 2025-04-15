package models

type Player struct {
	CurrentPlayer string `json:"currentPlayer"`
	Row           int    `json:"row"`
	Col           int    `json:"col"`
}
