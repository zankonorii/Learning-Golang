package main

import (
	"fmt"
	"os"
)

/*
 * Modify the echo program to print the index and value of each of its arguments, one per line.
 */

func main() {
	for key, value := range os.Args[1:] {
		fmt.Printf("%d:\t%s\n", key, value)
	}
}
