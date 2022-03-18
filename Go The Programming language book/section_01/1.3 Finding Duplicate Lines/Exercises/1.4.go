package main

import (
	"fmt"
	"os"
)

/*
 * Modify dup2 to print the names of all files in whicheach duplicated line occurs.
 */

func main() {

	counts := make(map[string]int)
	files := os.Args[1:]

	if len(files) == 0 {
		fmt.Println("No files to process")
		return
	} else {
		for _, arg := range files {
			_, err := os.Open(arg)

			if err != nil {
				fmt.Fprintf(os.Stderr, "%v\n", err)
				continue
			}
			counts[arg]++
		}
	}

	for name, count := range counts {
		if count > 1 {
			fmt.Printf("%d\t%s\n", count, name)
		}
	}

	// run program with "go run 1.3.go test test test" command
}
