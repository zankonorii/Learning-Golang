package main

import "fmt"

func main() {
	for i := 0; i < 4; i++ {
		fmt.Println(i)
	}

	fmt.Println("=== === === === === === === === === === === === === ===")
	for i := 0; i < 4; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}
	fmt.Println("=== === === === === === === === === === === === === ===")
	for i := 0; i < 4; i++ {
		if i == 3 {
			break
		}
		fmt.Println(i)
	}
	fmt.Println("=== === === === === === === === === === === === === ===")

	a := 0

	for a < 3 {
		fmt.Println(a)
		a++
	}

	fmt.Println("=== === === === === === === === === === === === === ===")

	/*x := 5

	switch x {
	case x < 10:
		fmt.Println("X is smaller than 10")
	case x < 100:
		fmt.Println("X is smaller than 100")
	}*/
}
