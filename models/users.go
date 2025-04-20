package models

type User struct {
	UserID int
	Name   string
}

var Users = []User{
	{UserID: 1, Name: "Het"},
	{UserID: 2, Name: "Rushil"},
}
