package main

import "fmt"

/*
 *	2. Using a Composite Literal :
 *		# create a Slice of type int.
 *		# assign 10 values
 *	 . Range over the slice and print the values out.
 *	 . Using format printing
 *	 	# print out the Type of the slice
 */

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	for index, value := range x {
		fmt.Println("index of ", index)
		fmt.Println("value is ", value)
	}
	fmt.Printf("%T\n", x)
}
