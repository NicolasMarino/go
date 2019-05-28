package models

type RestoSoft struct {
	Date     string  `json:"date"`
	Notes    string  `json:"notes"`
	Total    float32 `json:"total"`
	Items    []ItemsRs
	Customer Customer
	Business Restaurant
}
