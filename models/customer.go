package models

type Customer struct {
	ID       int64            `json:"id,omitempty" xml:"id,omitempty"`
	Name     string           `json:"name,omitempty" xml:"name"`
	LastName string           `json:"lastname,omitempty" xml:"lastname"`
	Email    string           `json:"email,omitempty" xml:"email,omitempty"`
	Location AddressRestoSoft `json:"location" xml:"location,omitempty"`
}

func (customer *Customer) GetFullName() string {
	return customer.Name + " " + customer.LastName
}
