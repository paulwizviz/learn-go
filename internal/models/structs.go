package models

type Simple struct {
	FirstName string
	Surname   string
}

type MultiTag struct {
	FirstName string `json:"firstname" c:"first_name"`
	Surname   string `json:"surname" c:"surname"`
}
