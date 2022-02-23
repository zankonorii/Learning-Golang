package main

import "fmt"

/*
 *	1. Create your own type "person" which will have an underlying type of "struct" so that it can store
 *	the following data :
 *		#	first name
 *		#	last name
 *		#	favorite ice cream flavors
 *	Create two Values of type person . Print out the values, ranging over tha elements in the slice.
 */

type person struct {
	first       string
	last        string
	favFavorite []string
}

func main() {
	p1 := person{
		first:       "James",
		last:        "Bond",
		favFavorite: []string{"chocolate", "martini", "rum and coke"},
	}

	p2 := person{
		first:       "Miss",
		last:        "moneyPenny",
		favFavorite: []string{"strawberry", "vanilla", "capuccino"},
	}

	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favFavorite {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favFavorite {
		fmt.Println(i, v)
	}

}
