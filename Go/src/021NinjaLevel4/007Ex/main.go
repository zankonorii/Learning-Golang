package main

import "fmt"

/*
 *	7.	Create a slice of string ([][]string). Store the following data in the multi-dimensional slice :
 *		# "James", "Bond", "Shaken, not stirred"
 *		# "Miss", "MoneyPenny", "Helloooooo, James."
 *	Range over the records, then range over the data in each record.
 *
 */

func main() {
	x := [][]string{
		{"James", "Bond", "Shaken, not stirred"},
		{"Miss", "MoneyPenny", "Helloooooo, James."},
	}
	fmt.Println(x)

	//Range Row
	for index, value := range x {
		fmt.Println(index, value)
	}

	//Range Data
	for index, value := range x {
		for ind2, data := range value {
			fmt.Println(index, ind2, data)
		}
	}
}
