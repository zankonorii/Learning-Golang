package main

import "fmt"

/*
 *	8.	Create a program that a switch statement with no switch expression specified.
 */

func main() {
	x := 1

	switch x {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	default:
		fmt.Println("I dont know this number")
	}

}
