package models

type Item struct {
	ID        int64     `json:"id"`
	Total     string	   `json:"total"`
	Quantity   string    `json:"quantity"`
	Name    string    `json:"name"`
	Options []Option `json:"options"`
}