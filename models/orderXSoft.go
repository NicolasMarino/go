package models

type Order struct {
	Customer Customer         `xml:"customer"`
	Business Restaurant       `xml:"business"`
	Status   string           `xml:"status"`
	Date     DateXResto       `xml:"date"`
	Notes    string           `xml:"notes"`
	Total    float32          `xml:"total"`
	Items    []ItemsRestoSoft `xml:"item"`
}
