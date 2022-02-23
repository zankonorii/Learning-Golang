package main

import "fmt"

func main() {
	//x := []type{values}	composite literal
	/*
	 *	# a slice allows you to group together Values of the same Type
	 */
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	fmt.Println(x)

	for index, value := range x {
		fmt.Println(index, value)
	}

	/*
	 *	[from to or equal that : up to Not equal that]
	 */
	fmt.Println(x[0:3])

	//append to s slice
	x = append(x, 10, 11, 12)
	fmt.Println("After appending ", x)

	y := []int{55, 66, 77, 88}
	fmt.Println("Y is ", y)

	/*
	 *	...T unlimited number of this type
	 *	y... mean all values of y
	 */
	x = append(x, y...)
	fmt.Println(x)

	//Delete from slice

	x = append(x[:2], x[6:]...)
	//OR x = append(y[:2], y[4:]...)
	fmt.Println(x)

	//Make function
	z := make([]int, 10, 100)
	fmt.Println(z)
	fmt.Println(len(z)) //return 10
	fmt.Println(cap(z)) //return 100

	//z[10] = 15 return an error
	z = append(z, 123) //now added new index to array
	fmt.Println(z)
	fmt.Println(len(z)) //return 11
	fmt.Println(cap(z)) //return 100
	//fmt.Println(z[9])

	/*
	 *	append add to Z until that our flow after that , max length (100) will increase automatic.
	 */

	//Multi Dimensional slice
	jb := []string{"James", "Bond", "Chocolate", "martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Moneypenny", "Strawberry", "Hazelnut"}
	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)
}
