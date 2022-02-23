package main

import "fmt"

const (
	_  = iota
	KB = 1 << (iota * 10)
	MB = 1 << (iota * 10)
	GB = 1 << (iota * 10)
)

func main() {
	/*x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1
	fmt.Printf("%d\t\t%b\n", y, y)

	z := x >> 1
	fmt.Printf("%d\t\t%b\n", z, z)*/
	/*B := 1024
	KB := 1024 * B
	MB := 1024 * KB
	GB := 1024 * MB
	TB := 1024 * GB

	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)*/

	/*B := 1024
	KB := B << 10
	MB := KB << 10
	GB := MB << 10
	TB := GB << 10

	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)
	fmt.Println(TB)*/

	fmt.Println(KB)
	fmt.Println(MB)
	fmt.Println(GB)

}
