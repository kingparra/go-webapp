package main

import "fmt"

func main() {
	// ///////////////////////////////
	// for i := 0; i <= 5; i++ {
	// fmt.Println("i is", i)
	// }
	// ////////////////////////////////
	// animals := []string{"dog", "fish", "horse", "cat"}
	// use the _ identifier to discard a value
	// for i, animal := range animals {
	// 	fmt.Println(i, animal)
	// }
	// ////////////////////////////////
	// animals := make(map[string]string)
	// animals["dog"] = "Fido"
	// animals["cat"] = "Fluffy"
	// for animalType, animal := range animals {
	// 	fmt.Println(animalType, animal)
	// }
	// ///////////////////////////////
	// firstLine := "Once upon a midnight dreary"
	// for i, l := range firstLine {
	// 	log.Println(i, ":", l) // print the index and the runes byte value
	// }
	// ///////////////////////////////
	// // Ranging over a struct
	type User struct {
		FirstName string
		LastName  string
		Email     string
		Age       int
	}
	users := []User{}
	// append : []type{} -> []type{} -> []type{}
	users = append(users, User{"John", "Smith", "john@smith.com", 30})
	users = append(users, User{"Mary", "Jones", "mary@jones.com", 30})
	for _, user := range users {
		fmt.Println(user.FirstName, user.LastName, user.Age)
	}
}
