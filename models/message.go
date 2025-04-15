package models

type Message struct {
	NextPlayer string `json:"nextPlayer"`
	Message    string `json:"message"`
}
