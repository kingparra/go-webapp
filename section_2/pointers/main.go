package main

import "log"

func main() {
	var myString string
	myString = "Green"
	log.Println("myString is set to", myString)
	// Create a reference with &
	changeUsingPointer(&myString)
}

// *string is type "pointer to string"
func changeUsingPointer(s *string) {
	newValue := "Red"
	// Dereference with *. For the variable at reference stored in s,
	// assign newValue.
	*s = newValue
}
