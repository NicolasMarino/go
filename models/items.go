package models

type Item struct {
	ID       int64    `json:"id,omitempty"`
	Name     string   `json:"name,omitempty"`
	Quantity float32  `json:"quantity"`
	Total    float32  `json:"total"`
	Options  []Option `json:"options,omitempty"`
}
