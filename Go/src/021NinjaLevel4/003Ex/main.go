package main

import "fmt"

/*
 *	3. Using the Code from the previous example, use Slicing to create the
 *		following new slices
 *	 Which are then printed :
 *	 	# [42 43 44 45 46]
 *	 	# [47 48 49 50 51]
 *	 	# [44 45 46 47 48]
 *	 	# [43 44 45 46 47]
 */

func main() {
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Println(x[:5])
	fmt.Println(x[5:])
	fmt.Println(x[2:7])
	fmt.Println(x[1:6])
}
