package models

type Todo struct {
	ID           int64  `json:"id"`
	Title        string `json:"title"`
	Desctription string `json:"description"`
	Done         bool   `json:"done"`
}
