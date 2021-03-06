package models

type Customer struct {
	ID       int64            `json:"id,omitempty" xml:"id,omitempty"`
	Name     string           `json:"name,omitempty" xml:"name,omitempty"`
	LastName string           `json:"lastname,omitempty" xml:"lastname,omitempty"`
	Email    string           `json:"email,omitempty" xml:"email,omitempty"`
	Location AddressRestoSoft `json:"location,omitempty" xml:"location,omitempty"`
}

func (customer *Customer) GetFullName() string {
	return customer.Name + " " + customer.LastName
}
