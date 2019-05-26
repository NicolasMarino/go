package models

type Option struct {
	ID       int64   `json:"id"`
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
}
