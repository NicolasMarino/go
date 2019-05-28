package models

type Restaurant struct {
	ID   int64  `json:"id,omitempty" xml:"id,omitempty"`
	Name string `json:"name" xml:"name"`
}

func (r Restaurant) getName() string {
	return r.Name
}
