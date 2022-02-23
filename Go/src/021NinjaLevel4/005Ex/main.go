package main

import "fmt"

/*
 *	5.	TO delete from a slice, we use Append along with slicing.
 *		For this hands-on exercise, follow these steps :
 *			# start with this slice
 *				x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
 *			# use APPEND $ SLICING to get the values here which you should then print:
 *				[42, 43, 44, 48, 49, 50, 51]
 */

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	y := []int{}
	y = append(x[:3], x[6:]...)
	//y = append(x[6:])
	fmt.Println(x)
	fmt.Println(y)
}
