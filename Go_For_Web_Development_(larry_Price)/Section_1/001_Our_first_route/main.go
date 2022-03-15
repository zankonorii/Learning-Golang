package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Hello web develpment")

	// this func get two arguments, route mach string and a function  to handel the request
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go web develpment")
	})

	fmt.Println(http.ListenAndServe("localhost:8080", nil))
}
