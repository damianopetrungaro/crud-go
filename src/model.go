package main

type User struct {
	Id        string
	FirstName string
	LastName  string
	Address   Address
}

type Address struct {
	City  string
	State string
}
