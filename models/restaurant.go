package models

type Restaurant struct {
	ID   int64  `json:"id,omitempty"`
	Name string `json:"name"`
}

func (r Restaurant) getName() string {
	return r.Name
}
