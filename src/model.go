package main

type User struct {
	Id        string
	Firstname string
	Lastname  string
	Address   Address
}

type Address struct {
	City  string
	State string
}
