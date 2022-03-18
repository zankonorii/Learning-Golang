package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

/*
 * Experiment  to measure the difference in running time between our potebtially inefficient
 * versions and the one that uses strings.Join.
 */

func main() {

	start := time.Now()
	var sep string

	for _, arg := range os.Args[1:] {
		sep += arg + " "
	}

	fmt.Println(sep)

	elapsed := time.Since(start)
	fmt.Printf("My Code %s\n", elapsed)

	start = time.Now()

	fmt.Println(strings.Join(os.Args[1:], " "))
	elapsed = time.Since(start)
	fmt.Printf("Using Join %s\n", elapsed)

}
