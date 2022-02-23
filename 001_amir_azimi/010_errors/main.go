package main

import (
	"fmt"
	"math"
)

func sqrt(n float64) (float64, error) {
	if n < 0 { //if negative
		//Errorf	== throw
		return 0.0, fmt.Errorf("sqrt of negative value (%f)", n)
	}

	return math.Sqrt(n), nil //	nil == null => not a number in JS
}
func main() {
	number, err := sqrt(64)

	//in go there aren't ===
	if err != nil {
		fmt.Printf("Error : %s\n", err)
	} else {
		fmt.Println(number)
	}

	number, err = sqrt(-64)
	if err != nil {
		fmt.Printf("Error : %s\n", err)
	} else {
		fmt.Println(number)
	}

	number, err = sqrt(64)
	fmt.Println(number, err)

	number, err = sqrt(-64)
	fmt.Println(number, err)

	//Using of panic is discouraged in GO.
}
