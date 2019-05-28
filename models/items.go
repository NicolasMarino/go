package models

type Item struct {
	ID       int64    `json:"id"`
	Name     string   `json:"name"`
	Quantity float32  `json:"quantity"`
	Price    float32  `json:"total"`
	Options  []Option `json:"options"`
}
