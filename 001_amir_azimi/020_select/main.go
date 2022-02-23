package main

import (
	"fmt"
	"time"
)

/*
 *	Select is like switch but it is for channels
 */

func main() {
	ch1, ch2 := make(chan int), make(chan int)

	go func() {
		ch1 <- 42
	}()

	//if channel is from ch1 => print First line
	//if channel is from ch2 => print Second line
	select {
	case val := <-ch1:
		fmt.Printf("Gote %d from ch1\n", val)
	case val := <-ch2:
		fmt.Printf("Gote %d from ch2\n", val)
	}

	fmt.Println("=== === === === === === === === === === === ===")

	out := make(chan float64)

	go func() {
		time.Sleep(100 * time.Millisecond)
		out <- 3.14
	}()

	select {
	case value := <-out:
		fmt.Printf("Gote %f \n", value)
	case <-time.After(200 * time.Millisecond): //should use <- because select is only for channeles
		fmt.Println("time out")

	}
	/*
		چون بالایی تایمش کمتر است (100) پس case اول اجرا می شود.
		اگر تایم اول را برای مثال از 100 به 300 تغییر بدهیم case دوم اچرا می شود.
	*/
}
