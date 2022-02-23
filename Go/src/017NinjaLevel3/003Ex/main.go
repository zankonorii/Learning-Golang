package main

import "fmt"

/*
 * 	3. Create a for loop using this syntax
 * 		# for condition {}
 * 	Have it print out the years you been alive
 */

func main() {
	bd := 1999

	for bd <= 2021 {
		fmt.Println(bd)
		bd++
	}
}
