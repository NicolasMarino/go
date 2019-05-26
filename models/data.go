package models

import (
	"time"
)

// Represents the order
type Data struct {
	ID	int64	`json:"id"`
	State	string	`json:"state"`
	Pickup	bool	`json:"pickup"`
	Notes	string	`json:"notes"`
	RegisteredDate time.Time `json:"registeredDate"`
	Integration string	`json:"integration"`
	Customer Customer	`json:"customer"`
	Address	Address	`json:"address"`
	Restaurant	Restaurant	`json:"restaurant"`
	Total	string	`json:"total"`
	Shipping	string	`json:"shipping"`
	Subtotal	string    `json:"subtotal"`
	Items []Item	`json:"items"`
}
