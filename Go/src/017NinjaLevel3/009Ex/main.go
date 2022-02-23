package main

import "fmt"

/*
 *	9. Create a Program that uses a switch statement with the switch expression
 *		specified as a variable of TYPE string the Identifier "favSport"
 */

func main() {
	favSport := "surfing"

	switch favSport {
	case "skiing":
		fmt.Println("go the mountains!")
	case "swimming":
		fmt.Println("go the pool!")
	case "surfing":
		fmt.Println("go to hawaii!")
	case "wingsuit flying":
		fmt.Println("what would you like to say at your funeral?")
	}

}
