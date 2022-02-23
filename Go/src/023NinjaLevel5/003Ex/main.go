package main

import "fmt"

/*
 *	3.	Create a new Type: vehicle.
 *		#	The underlying type is a struct.
 *		#	The fields :
 *				->	doors
 *				->	color
 *	Create two new types: truck & sedan.
 *		#	The underlying type of each these new types is a struct.
 *		#	Embed The "vehicle" type in both truck and sedan.
 *		#	Give Truck the field "fourWheel" which will be set to bool.
 *		#	Give sedan the field "luxury" which will be set to bool.
 *	Using the vehicle, truck, and sedan structs:
 *		# Using a composite literal, create a value of type truck and assign
 *		values to the fields.
 *		#	Using Composite literal, create a value of type sedan and assign
 *		values to the fields.
 *		#	Print out each of these values.
 *		#	Print out a single field
 */

type vehicle struct {
	doors int
	color string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t := truck{
		vehicle: vehicle{
			doors: 2,
			color: "white",
		},
		fourWheel: true,
	}
	fmt.Println("Truck is ", t)

	s := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "black",
		},
		luxury: false,
	}
	fmt.Println("Sedan is ", s)

	fmt.Println(t.doors)
	fmt.Println(t.vehicle.doors)
	fmt.Println(t.color)
	fmt.Println(t.vehicle.color)
	fmt.Println(t.fourWheel)

	fmt.Println(s.doors)
	fmt.Println(s.vehicle.doors)
	fmt.Println(s.color)
	fmt.Println(s.vehicle.color)
	fmt.Println(s.luxury)
}
