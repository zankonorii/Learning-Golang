package main

import (
	"fmt"
)


func HelloCodeCup(n int) string {
	return "Hello CodeCup " + fmt.Sprintf("%d", n)
}

func main() {
	str := HelloCodeCup(6)
	fmt.Println(str)
}
