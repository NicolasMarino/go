package models

type Option struct {
	ID       int64   `json:"id,omitempty"`
	Name     string  `json:"name,omitempty"`
	Quantity float32 `json:"quantity,omitempty"`
}
