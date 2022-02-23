package main

import "fmt"

func main() {
	x := 2

	/*switch x {
	case 1:
		fmt.Println("X is equal to One")
	case 2:
		fmt.Println("X is equal to Two")
	case 3:
		fmt.Println("X is equal to Three")
		fallthrough ///*****
	default:
		fmt.Println("X is an unknown number")
	}*/

	//multi cases
	switch x {
	case 1, 2, 3:
		fmt.Println("X is on of the 1,2,3")
		fmt.Println("this is another sentence")
		fmt.Println("this is another sentence")
		fmt.Println("this is another sentence")
	case 4, 5, 6:
		fmt.Println("X is one of the 4,5,6")
	case 7, 8, 9:
		fmt.Println("X is one of the 7,8,9")
	default:
		fmt.Println("X is an unknown number")
	}
}
