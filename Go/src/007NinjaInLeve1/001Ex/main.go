package main

import "fmt"

/*
 * 1. 	Using the short declaration operator, Assign these Values
 * 		to variables with the Identifier "x", "y" and "z"
 *			a. 	42
 *			b. 	"Jams Bond"
 *			c. 	true
 *
 * 2. 	Now print the values stored in those variables using
 *			a.	a single print statement
 *			b.	multiple print statements
 */
func main() {
	//answer of 1
	x := 42
	y := "Jams Bond"
	z := true

	//answer 2.a
	fmt.Println(x, y, z)
	//answer 2.b
	fmt.Println("X is : ", x)
	fmt.Println("Y is : ", y)
	fmt.Println("Z is : ", z)
}
