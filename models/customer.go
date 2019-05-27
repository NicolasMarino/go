package models

type Customer struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Email    string `json:"email"`
	Adress   Address
}
