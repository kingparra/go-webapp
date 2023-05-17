package main

// first line must be a package declaration

import "fmt"

func main() {
	fmt.Println("Hello, world.")
	var whatToSay string = "Goodbye, cruel world."
	var i int
	fmt.Println(whatToSay)
	i = 7
	fmt.Println("i is set to", i)
	whatWasSaid, theOtherThingThatWasSaid := saySomething()
	fmt.Println("The function returned", whatWasSaid, theOtherThingThatWasSaid)
}

func saySomething() (string, string) {
	return "something", "else"
}
