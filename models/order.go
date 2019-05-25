package models

import (
	"time"
)

// Represents the order
type Order struct {
	ID        int64     `json:"id"`
	State     bool	   `json:"state" validate:"required"`
	Pickup   string    `json:"pickup" validate:"required"`
	Notes    string    `json:"notes"`
	RegisteredDate time.Time `json:"registeredDate"`
	Integration string	`json:"integration"`
	Customer Customer	 `json:"customer"`
	Address   Address    `json:"address"`
	Restaurant    Restaurant	`json:"restaurant"`
	Total    string    `json:"total"`
	Shipping    string    `json:"shipping"`
	Subtotal    string    `json:"subtotal"`
	Items []Item  `json:"items"`
}
