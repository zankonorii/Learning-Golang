package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

type Job struct {
	User   string `json:"user"`
	Action string `json:"action"`
	Count  int    `json:"count"`
}

func main() {
	// Get Request
	response, err := http.Get("https://httpbin.org/get")

	if err != nil {
		log.Fatalf("ERROR : can't Get %s\n", err)
	}

	defer response.Body.Close() // If We don't close it, it gets memory from over OS

	io.Copy(os.Stdout, response.Body)
	//fmt.Printf("Get is %+v\n", response)

	// === === === === === === === === === === === === === === === === === === === === === === === === === === === ===

	// Post Request
	job := &Job{
		User:   "Amir",
		Action: "punch",
		Count:  1,
	}

	var buf bytes.Buffer // Buffer is a memory reader => in-memory reader or io reader

	enc := json.NewEncoder(&buf) // Like a constructor

	if err := enc.Encode(job); err != nil {
		log.Fatalf("ERROR : can't encode job - %s\n", err)
	}

	response, err = http.Post("https://httpbin.org/post", "application/json", &buf)

	if err != nil {
		log.Fatalf("ERROR : can't call Post method %s\n", err)
	}

	defer response.Body.Close()

	io.Copy(os.Stdout, response.Body)
}
