package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fprintf, err := fmt.Fprintf(w, "hello parsclick fans!")
	if err != nil {
		fmt.Printf("ERROR : %s\n", err)
	}

	fmt.Println(fprintf)
}

func main() {
	// http.HandleFunc is like route in laravel
	// http.HandleFunc(url, function) Note : is call function (like this func())
	// function should return a value else we can just run it without any ()
	http.HandleFunc("/hello", helloHandler)

	// ListenAndServe isa function that listen to a port
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
