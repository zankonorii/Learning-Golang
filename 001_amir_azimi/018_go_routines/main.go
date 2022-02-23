package main

import (
	"fmt"
	"net/http"
	"sync"
)

func returntype(url string) {
	response, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}

	//buf := new(bytes.Buffer)
	//buf.ReadFrom(response.Header)
	//newStr := buf.String()
	defer response.Body.Close()
	contentType := response.Header.Get("Content-Type")
	//fmt.Printf("%s ->  %s\n", url, newStr)
	fmt.Printf("%s ->  %s\n", url, contentType)
}

func main() {
	urls := []string{
		"https://golang.org",
		"https://api.github.com",
		"https://httpbin.org/xml",
	}

	for _, url := range urls {

		//this section doesn't wait for `return type` func response
		go func(url string) {
			returntype(url)
		}(url) //url passed to anonymous function
	}

	//this section waits with our control

	//follow control steps : 1, 2 , ...
	var waitGroup sync.WaitGroup //2
	for _, url := range urls {
		waitGroup.Add(1) // waiting 1 delta // 3 in this step waiting until Done
		go func(url string) {
			returntype(url)
			waitGroup.Done() //4
		}(url) //url passed to anonymous function
	}
	waitGroup.Wait() //1
}
