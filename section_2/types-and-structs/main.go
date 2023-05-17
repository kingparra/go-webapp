package main

import (
	"log"
	"time"
)

// Defining a simple struct, for use as a record.
type User struct {
	// Why caps? To export the fields. Only vars beginning with a capital letter are exported.
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	Birthdate   time.Time
}

// Associating a function with a struct type.
type myStruct struct {
	FirstName string
}

func (m *myStruct) printFirstName() string {
	return m.FirstName
}

func main() {
	user := User{
		FirstName:   "Trevor",
		LastName:    "Sawler",
		PhoneNumber: "1 555 555-1212",
	}
	log.Println(user.FirstName, user.LastName, user.Birthdate)
}
