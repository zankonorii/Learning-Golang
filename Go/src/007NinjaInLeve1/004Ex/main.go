package main

import "fmt"

/*
 * 1. Create your own type, Have the underlying type be an int.
 * 2. Create a Variable of your new type with identifier `x` using the `VAR` keyword
 * 3. in func main
 *	a.	print out value of the variable "X"
 *	b.	print out the type of variable "X"
 *	c.	assign 42 to the variable "x" using the `=` operator
 *	d.	print out the value of the variable "X"
 */

type MyType int

var x MyType

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
