package main

import "fmt"

func main() {
	foo := []string{"amir", "azimi", "parsclick"}

	fmt.Printf("foo is = %v (type %T)\n", foo, foo)

	fooNumbers := []int{1, 2, 3, 4}

	fmt.Printf("foo numbers is %v (type %T)\n", fooNumbers, fooNumbers)

	fmt.Println("=== === === === === === === === === === === === === === === === === === === ===")

	//length
	fmt.Println("foo length is ", len(foo))
	fmt.Println("foo number length is ", len(fooNumbers))

	fmt.Println("=== === === === === === === === === === === === === === === === === === === ===")

	//0 base or 0 index
	fmt.Println("foo first index is ", foo[0])
	fmt.Println("foo number first index is ", fooNumbers[0])

	fmt.Println("=== === === === === === === === === === === === === === === === === === === ===")

	//slices
	fmt.Println("foo[1:]	=>	", foo[1:])
	fmt.Println("foo[:1]	=>	", foo[:1])

	fmt.Println("=== === === === === === === === === === === === === === === === === === === ===")

	//For loop
	for i := 0; i < len(foo); i++ {
		fmt.Println(foo[i])
	}

	fmt.Println("=== === === === === === === === === === === === === === === === === === === ===")

	//Single value range (to show indexes only)
	for index := range foo {
		fmt.Println(index)
	}

	//Double value range (to show index and value)
	for index, value := range foo {
		fmt.Println(index, "\t", value)
	}

	//Double value range (to show index and value)
	for index, value := range foo {
		//S => string , d => digit
		fmt.Printf("%s at %d\n", value, index)
	}

	//ignore index by using under score _
	for _, value := range foo {
		fmt.Println(value)
	}

	fmt.Println("=== === === === === === === === === === === === === === === === === === === ===")

	//Append
	foo = append(foo, "London")
	fmt.Println(foo)
}
