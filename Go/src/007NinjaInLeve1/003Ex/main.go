package main

import "fmt"

/*
 * Using the code from previous exercise,
 *		1. At the package level scope, Assign the following values to three variables
 *			a. for `x` assign 42
 *			b. for `y` assign "James Bond"
 *			c. for `z` assign true
 *
 *		2. in func main
 *			a. use fmt.Sprintf to print all of the Values to one single string.
 *				Assign the returned value of Type string using the short declaration
 *				operator to a Variable with the Identifier `s`
 *			b. print out the value stored by variable `s`
 */

//Answer 1
var x int = 42
var y string = "Jams Bond"
var z bool = true

func main() {
	//answer 2.a
	s := fmt.Sprintf("%v,\t%v,\t%v", x, y, z)

	//answer 2.b
	fmt.Println(s)
}
