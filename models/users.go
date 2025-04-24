package models

type User struct {
	UserID int
	Name   string
}

var Users = []User{
	{UserID: 1, Name: "Het"},
	{UserID: 2, Name: "Vraj"},
	{UserID: 3, Name: "Fenil"},
	{UserID: 4, Name: "Nemi"},
	{UserID: 5, Name: "Rushil"},
}
