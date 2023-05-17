package main

import "log"

type User struct {
	FirstName string
	LastName  string
}

func main() {
	// //////////////////////////////////////////
	// //                 V key  V value
	// myMap := make(map[string]string)

	// // Read up on the builtin.make and builtin.new functions

	// myMap["dog"] = "Samson"
	// myMap["other-dog"] = "Cassie"
	// // replacing a value
	// myMap["dog"] = "Fido"

	// log.Println(myMap["dog"])
	// log.Println(myMap["other-dog"])

	// //////////////////////////////////////////
	// myMap := make(map[string]int)

	// myMap["First"] = 1
	// myMap["Second"] = 2
	// log.Println(myMap["First"], myMap["Second"])

	// //////////////////////////////////////////
	// myMap := make(map[string]User)
	// me := User{
	// 	FirstName: "Trevor",
	// 	LastName:  "Sawler",
	// }

	// myMap["me"] = me

	// log.Println(myMap["me"].FirstName)

	// //////////////////////////////////////////
	// Maps are immutable!
	// Maps are unordered, you must always look up by key.

	// Say you want to create a map, but don't know what the type of the value should be. You can use interface{}.
	// myMap := make(map[string]interface{})

	// //////////////////////////////////////////
	// // Slices! They're like arrays, sort of.
	// var mySlice []int
	// // [2 1 3]
	// mySlice = append(mySlice, 2)
	// mySlice = append(mySlice, 1)
	// mySlice = append(mySlice, 3)

	// // [1 2 3]
	// sort.Ints(mySlice)

	// log.Println(mySlice)

	// ///////////////////////////////////////////
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// slicing syntax is like Python
	log.Println(numbers[0:2])
}
