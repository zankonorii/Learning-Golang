package main

import "fmt"

/*
 * #5. Create a Variable of Type string using a raw string literal. Print it.
 */

func main() {
	x := `this is a string for exercise 5
	  raw string "string"`
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}
