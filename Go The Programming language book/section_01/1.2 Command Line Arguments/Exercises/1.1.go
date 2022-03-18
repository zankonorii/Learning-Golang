package main

import (
	"fmt"
	"os"
	"strings"
)

/*
 * Modify the echo program to also print os.Args[0], the name of the command that invoked it.
 */

func main() {
	var args string

	firstArgument := strings.Split(os.Args[0], "/")
	args = firstArgument[len(firstArgument)-1] + " "

	for _, arg := range os.Args[1:] {
		args += arg + " "
	}

	fmt.Println(args)
}
