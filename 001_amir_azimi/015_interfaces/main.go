/*
 *	Interface is a protocol that when we use it, every struct should hae here methods
 */
package main

import (
	"fmt"
	"math"
)

type Square struct {
	Length float64
}

type Circle struct {
	Radios float64
}

type Shape interface {
	Area() float64
}

func (square *Square) Area() float64 {
	return square.Length * square.Length
}

func (circle *Circle) Area() float64 {
	return math.Pi * math.Pow(circle.Radios, 2)
}

func sumArea(shapes []Shape) float64 { // input slice => Note : slice = array
	total := 0.0
	for _, shape := range shapes {
		total += shape.Area()
	}

	return total
}
func main() {
	//or ypu can : &square := &Square{5}
	square := Square{5}
	fmt.Println(square)
	fmt.Println(square.Area())

	//or you can : &circle := Circle{3}
	circle := Circle{3}
	fmt.Println(circle)
	fmt.Println(circle.Area())

	//shapes := []Shape{square, circle}  // if in declaring use & in there not need it.
	shapes := []Shape{&square, &circle}
	//fmt.Println(sumArea([]Shape{&square, &circle})) // this is another way
	fmt.Println(sumArea(shapes))
}
