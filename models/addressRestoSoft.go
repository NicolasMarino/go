package models

type AddressRestoSoft struct {
	Longitude   string `json:"latitude" xml:"latitude,omitempty"`
	Latitude    string `json:"longitude" xml:"longitude,omitempty"`
	Coordinates string `json:"coordinates,omitempty" xml:"coordinates"`
}
