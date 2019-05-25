package models

type Item struct {
	ID        int64     `json:"id"`
	Total     bool	   `json:"state" validate:"required"`
	Quantity   string    `json:"pickup" validate:"required"`
	Name    string    `json:"notes"`
	Options []Option `json:"Option"`
}