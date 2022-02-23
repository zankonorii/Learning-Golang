package main

import (
	"fmt"
	"time"
)

/*
 *  We have two kins of channels : (Buffered, UnBuffered)
 *
 *  === === === === === === === === === === === ===
 *  [ | | ] => this is a channel like a queue
 *  === === === === === === === === === === === ===
 *
 *  UnBuffered :
 * 		if you send a value to this channel we blocked
 * 	until send this data
 * 	[1| | ]		=>	insert data, we blocked
 * 	[ |1| ]		=>	move   data, we blocked yet
 * 	[ | |1]		=>	move   data, we blocked yet
 * 	[ | | ]		=>	data resided, we unblocked and channel opened
 *  === === === === === === === === === === === ===
 *
 *	Buffered:
 *		just when channel is full blacked like follow EX
 *
 *	=== === === === === === === === === === === ===
 *  BUFFERED CHANNEL EXAMPLE (capacity : 3)
 *	=== === === === === === === === === === === ===
 *	NOT BLocked	->	[ |	| ]
 *	NOT BLocked	->	[1|	| ]
 *	NOT BLocked	->	[1|2| ]
 *	BLocked		->	[1|2|3]
 * 	=== === === === === === === === === === === ===
 */

func main() {
	//make create a channel .
	channel := make(chan int)

	/*//This will block
	//use channels with <- operator
	<-channel
	//This will block because nothing was pushed to the channel
	fmt.Println("here")
	// This section return an error because we don't use channel*/

	//go func + tab to create a go routine
	go func() {
		//send number of the channel
		channel <- 350 // send 350 to channel
	}()

	//Receive number from the channel

	val := <-channel

	fmt.Printf("got %d\n", val)
	fmt.Println("=== === === === === === === === === ===")

	//Send Multiple
	/*
		//1 ):
			for i := 0; i < 3; i++ {
				fmt.Printf("Sending Digit %d\n", i)

				go func(i int) {
					channel <- i
				}(i)
			}
	*/
	//2 ):
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			channel <- i
			time.Sleep(1 * time.Second) //1 second
		}
	}()
	//Receive Multiple
	for i := 0; i < 3; i++ {
		value := <-channel
		fmt.Printf("Got Digit %d\n", value)
	}

	/*
			 * if we run above code that give us this output :
			 *		sending 0
			 *		Got Digit 0
			 *		sending 1
			 *		Got Digit 1
			 *		sending 2
			 *		Got Digit 2
		 	 *	that is an unbuffered channel
		 	 *
		 	 *	channels at default are unbuffered
		 	 *
		 	 *	if second for loop is less than first ,
		 	 *	the result just return second loop
		 	 *
	*/

	fmt.Println("=== === === === === === === === === === === ===")
	//Close to Signal we are done Multiple
	go func() {
		for i := 0; i < 5; i++ {
			fmt.Printf("Sending %d\n", i)
			channel <- i
			time.Sleep(time.Second)
		}
		close(channel) // this function close the channel
		//this function helps us for get unlimited channels
	}()

	for i := range channel {
		fmt.Printf("Recive %d\n", i)
	}
}
