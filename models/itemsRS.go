package models

type ItemsRs struct {
	Name     string  `json:"name"`
	Quantity float32 `json:"quantity"`
	Price    float32 `json:"price"`
}
