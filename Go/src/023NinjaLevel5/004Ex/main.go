package main

import "fmt"

/*
 *	4.	Create and Use an anonymous struct.
 */

func main() {
	/*
		// My Solution
		me := struct {
			name     string
			lastName string
			age      int
			avr      float32
			favorite []string
			works    map[string]string
		}{
			name:     "zanko",
			lastName: "nori",
			age:      23,
			favorite: []string{"code", "learning", "computer", "reading", "eating"},
			works: map[string]string{
				"birka":     "multi store",
				"azar park": "parking service",
				"alget":     "Multi store website",
				"samsung":   "learning education",
			},
		}

		fmt.Println(me)

		fmt.Println("Favorites   : ")
		for _, key := range me.favorite {
			fmt.Println(key)
		}

		for key, value := range me.works{
			fmt.Println("Title : ", key)
			fmt.Println("Description : ", value)
		}*/

	//Teacher Solution
	s := struct {
		first     string
		friends   map[string]int
		favDrinks []string
	}{
		first: "James",
		friends: map[string]int{
			"moneyPenny": 555,
			"Q":          777,
			"M":          888,
		},
		favDrinks: []string{
			"Martini",
			"Water",
		},
	}

	fmt.Println(s)

}
