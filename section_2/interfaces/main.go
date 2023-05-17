package main

import "fmt"

// Define the interface
type Animal interface {
	Says() string
	NumberOfLegs() int
}

// Create some types we intend to implement the interface for
type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

// Implement the class methods for the types
func (d *Dog) Says() string {
	return "woof"
}

func (d *Dog) NumberOfLegs() int {
	return 4
}

func (d *Gorilla) Says() string {
	return "Ugh"
}

func (d *Gorilla) NumberOfLegs() int {
	return 2
}

// This is a helper function to print info about the types
func PrintInfo(a Animal) {
	fmt.Println("This animal says", a.Says(), "and has", a.NumberOfLegs())
}

// ...and here we call it
func main() {
	dog := Dog{
		Name:  "Samson",
		Breed: "German Shepard",
	}
	PrintInfo(&dog)
	gorilla := Gorilla{
		Name:          "Jock",
		Color:         "grey",
		NumberOfTeeth: 38,
	}
	PrintInfo(&gorilla)
}
