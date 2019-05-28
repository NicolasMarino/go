package models

type Order struct {
	Customer Customer         `xml:"customer"`
	Business Restaurant       `xml:"business,omitempty"`
	Status   string           `xml:"status,omitempty"`
	Date     DateXResto       `xml:"date,omitempty"`
	Notes    string           `xml:"notes,omitempty"`
	Total    float32          `xml:"total,omitempty"`
	Items    []ItemsRestoSoft `json:"items,omitempty" xml:"items,omitempty"`
}
