package models

// Represents the order
type Orders struct {
	Count int64  `json:"count"`
	Datos []Data `json:"data"`
}
