package models

type Customer struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	LastName string `json:"lastname,omitempty"`
	Email    string `json:"email,omitempty"`
	Location AddressRestoSoft
}
