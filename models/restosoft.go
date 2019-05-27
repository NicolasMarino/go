package models

import (
	"time"
)

type RestoSoft struct {
	Date     time.Time `json:"date"`
	Notes    string    `json:"notes"`
	Total    float32   `json:"total"`
	Items    []Item    `json:"items"`
	Customer Customer  `json:"customer"`
	Business Restaurant
}

/*func (r RestoSoft) restoSoft() string {

	datos := r.Date, r.Notes, r.Total, r.Items, r.Customer, r.Business.getName
	return datos
}*/
