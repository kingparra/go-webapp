package main

import "log"

func main() {
	// ////////////////////////////////////////
	var isTrue bool
	isTrue = true

	// Go has the standard boolean opeators && || !
	if isTrue {
		log.Println("isTrue is", isTrue)
	} else {
		log.Println("isTrue is", isTrue)
	}
	// ////////////////////////////////////////
	myVar := "cat"
	switch myVar {
	case "cat":
		log.Println("cat is set to cat")
	case "dog":
		log.Println("cat is set to dog")
	default:
		log.Println("cat is something else")
	}
	// ///////////////////////////////////////
}
