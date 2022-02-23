package main

import "fmt"

/*
 *	7. Building on the previous hands-on exercise, create a program that uses "else if" and "else"
 */

func main() {
	x := "Money penny"

	if x == "Money penny" {
		fmt.Println(x)
	} else if x == "Jams Bond" {
		fmt.Println("BOND ::", x)
	} else {
		fmt.Println("unknown")
	}
}
