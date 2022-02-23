package main

import "fmt"

/*
 *	type of resources ::
 *		1) Memory
 *		2) Files
 *		3) Sockets
 *		4) Virtual Machines
 *
 */
func cleanup(name string) {
	fmt.Printf("Cleaning up %s\n", name)
}

func worker() {
	//It declared at first but called at last
	defer cleanup("A")
	defer cleanup("B")
		
	fmt.Println("Worker")
}

func main() {
	worker()
}
