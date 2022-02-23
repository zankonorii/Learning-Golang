package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var data = `
	{
		"user"   : "Amir Azimi",
		"type"	 : "deposit",
		"amount" : 100000.5
	}
`

// Request is a bank transaction
type Request struct {
	Login  string  `json:"user"`
	Type   string  `json:"type"`
	Amount float64 `json:"amount"`
}

func main() {
	rdr := bytes.NewBufferString(data) // Simulate a file/socket

	//Decode the request
	decode := json.NewDecoder(rdr) //like json decoder in php , or json parse in js

	request := &Request{} // Struct for a request

	if err := decode.Decode(request); err != nil {
		log.Fatalf("ERROR: can't decode - %s\n", err)
	}

	fmt.Printf("got : %+v\n", request)

	//How to Create response
	prevBalance := 850000000.0 // Loaded From Database

	response := map[string]interface{}{ //	Use map with empty Interface for response
		"ok":      true,
		"code":    200,
		"balance": prevBalance + request.Amount,
	}

	//Encode the response
	encode := json.NewEncoder(os.Stdout)

	if err := encode.Encode(response); err != nil {
		log.Fatalf("ERROR: can't encode %s\n", err)
	}

	fmt.Printf("got %+v\n", response)
}
