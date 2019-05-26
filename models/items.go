package models

type Item struct {
	ID       int64    `json:"id"`
	Total    float32  `json:"total"`
	Quantity float32  `json:"quantity"`
	Name     string   `json:"name"`
	Options  []Option `json:"options"`
}
