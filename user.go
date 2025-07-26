package main

type User struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name"`
}
