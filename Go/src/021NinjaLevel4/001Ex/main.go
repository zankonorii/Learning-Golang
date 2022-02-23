package main

import "fmt"

/*
 *	1. Using a Composite literal:
 *		# create an Array which holds 5 Values of type int
 *		# assign Values to each index position.
 *	? Range over the array and print the values out.
 *	? Using format printing
 *		# print out the Type of the array
 */

func main() {
	/*
		my result
		var x [5]int
		fmt.Println(x)

		x[0] = 1
		x[1] = 2
		x[2] = 3
		x[3] = 4
		x[4] = 5

		fmt.Println(x)

		for v, k := range x {
			fmt.Println("v  is : ", v)
			fmt.Println("k  is : ", k)
		}*/

	//teacher result
	x := [5]int{42, 43, 44, 45, 46}
	for index, value := range x {
		fmt.Println("index is : ", index)
		fmt.Println("value is : ", value)
	}

	//type of it
	fmt.Printf("%T\n", x)
}
