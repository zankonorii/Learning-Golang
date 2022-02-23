package main

import "fmt"

type Point struct {
	x int
	y int
}

// Move if we don't use * to reference that not be working currently
func (point Point) Move(dx, dy int) {
	point.x += dx
	point.y += dy
}

func (point *Point) MoveTow(dx, dy int) {
	point.x += dx
	point.y += dy
}

func main() {
	point := Point{1, 2}
	point.Move(2, 3)
	fmt.Printf("%+v\n", point)

	point.MoveTow(2, 3)
	fmt.Printf("%+v\n", point)
}
