package main

import "fmt"

/*
 *	Building on the code from the previous example
 *		1.	at the package level scope, using the `VAR` keyword, create a Variable
 *			with the Identifier "Y". The variable should be of the Underling type
 *			of your custom type "x"
 *		2.	In func main
 *			a.	this should already be done
 *				  i.	print out the value of the variable "x"
 *			   	 ii.	print out the type of the variable "x"
 *			  	iii.	assign your Value to the variable "x" using the `=` operator
 *			   	 iv.	print out the Value of the variable "x"
 *			b.	now do this
 *				i.	now use Conversion to convert the type of the value stored in "x"
 *					to the underlying type
 *						1.	then use the short declaration operator to assign that value
 *							to "y"
 *						2. print out the value stored in "Y"
 *						3. print out the type of "y"
 */

type MyType int

var x MyType
var y int

func main() {
	fmt.Println("value of x is : ", x)
	fmt.Printf("%T\n", x)

	x = 42

	fmt.Println("value of x after assign it : ", x)

	//Note : Go cannot convert string to int...
	y = int(x)

	fmt.Println("value of Y is : ", y)
	fmt.Printf("%T\n", y)
}
