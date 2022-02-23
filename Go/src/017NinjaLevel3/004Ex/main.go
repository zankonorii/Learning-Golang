package main

import "fmt"

/*
 *	4. Create a for loop using this syntax
 *		#	for{}
 *	Have it print out the years you have been alive.
 */

func main() {
	bd := 1999

	for {
		if bd > 2021 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
