package main

type User struct {
	UUID      string
	Firstname string
	Lastname  string
}

type Merchant struct {
	UUID string
	DBA  string
}

type Order struct {
	Amount   uint
	User     User
	Merchant Merchant
}
