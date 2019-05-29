package models

import "encoding/xml"

type Order struct {
	XMLName  xml.Name          `xml:"order"`
	Customer CustomerRestoSoft `xml:"customer"`
	Business Restaurant        `xml:"business,omitempty"`
	Status   string            `xml:"status,omitempty"`
	Date     DateXResto        `xml:"date,omitempty"`
	Notes    string            `xml:"notes,omitempty"`
	Total    float32           `xml:"total,omitempty"`
	Items    []ItemsRestoSoft  `xml:"items>item,omitempty"`
}
