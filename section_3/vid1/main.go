package main

import (
	"fmt"
	"net/http"
)

func main() {

	// What's going on with the parameters of this function?
	// What don't I understand about it, in the first place?
	// Is http.ResponseWriter a type or an interface...?
	// Is http.Request a type or interface?
	// Why is it a pointer to http.Request?
	//
	http.HandleFunc("/cruel", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Goodbye, cruel world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})
	http.HandleFunc("/kind", func(w http.ResponseWriter, r *http.Request) {
		n, err := fmt.Fprintf(w, "Hello, kind world!")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Bytes written: %d", n))
	})
	// start the http server
	_ = http.ListenAndServe(":8080", nil)
}

/*
Read the docs for: net/http, fmt.Sprintf
Read the wikipedia page on HTTP (the protocol)
Read about how to route http in go and Haskell
*/
