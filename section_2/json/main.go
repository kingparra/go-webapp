package main

import (
	"encoding/json"
	"fmt"
	"log"
)

// we can use this type to unmarshal the json blob into a struct
type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	HasDog    bool   `json:"has_dog"`
}

func main() {
	myJson := `
[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"has_dog": false
	}
]`
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("unmarshalled: %v", unmarshalled)

	// write json from a struct
	var mySlice []Person
	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HasDog = false

	mySlice = append(mySlice, m1)

	newJson, err := json.MarshalIndent(mySlice, "", "  ")
	if err != nil {
		log.Println("error marshalling", err)
	}
	fmt.Println(string(newJson))
}
