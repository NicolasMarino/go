package models

type Address struct {
	Description string `json:"description,omitempty"`
	Coordinates string `json:"coordinates,omitempty" `
	Phone       string `json:"phone,omitempty"`
}
