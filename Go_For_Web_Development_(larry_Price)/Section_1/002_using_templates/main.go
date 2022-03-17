package main

import (
	"fmt"
	"net/http"
    "html/template"
)

func main() {
    templetes := template.Must(template.ParseFiles("templates/index.htm"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello Zanko")
	})

	http.ListenAndServe(":8080", nil)
}
