package main

import "fmt"

func main() {
	/*for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Printf("%v\t", i*j)
		}
		fmt.Println()
	}*/

	for i := 33; i < 122; i++ {
		fmt.Printf("%#U\n", i)
	}
}
