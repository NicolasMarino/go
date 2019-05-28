package models

import (
	"time"
)

type Data struct {
	ID             int64      `json:"id"`
	State          string     `json:"state" xml:"state"`
	Pickup         bool       `json:"pickup"`
	Notes          string     `json:"notes"`
	RegisteredDate time.Time  `json:"registeredDate"`
	Integration    string     `json:"integration"`
	Customer       Customer   `json:"customer"`
	Address        Address    `json:"address"`
	Restaurant     Restaurant `json:"restaurant"`
	Total          float32    `json:"total"`
	Shipping       float32    `json:"shipping"`
	Subtotal       float32    `json:"subtotal"`
	Items          []Item     `json:"items"`
}
