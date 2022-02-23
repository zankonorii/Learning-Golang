package main

import "fmt"

/*
 * 1.Use `VAR` to Declare three variables. The Variables should have package
 *	level scope .
 * 	do not Assign Values to the Variables . Use The following Identifiers for
 *  the Variables and make sure the variables store values of the following type.
 * 		a. identifier "x" type int
 * 		b. identifier "y" type string
 * 		c. identifier "z" type bool
 *
 * 2. in func main
 *		a. print out the values for each identifier
 *		b. the compiler assigned values to the variables, what are these value called?
 */

var x int
var y string
var z bool

func main() {
	fmt.Println("value Of `x` is : ", x)
	fmt.Printf("%T\n", x)

	fmt.Println("value of `y` is : ", y)
	fmt.Printf("%T\n", y)
	fmt.Println("value of `z` is : ", z)
	fmt.Println("%T\n", z)

	//Answer 2
	fmt.Println("it called zero value")
}
