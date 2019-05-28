package models

type ItemsRestoSoft struct {
	Name     string  `json:"name" xml:"name"`
	Quantity float32 `json:"quantity" xml:"quantity"`
	Price    float32 `json:"price" xml:"price"`
}
