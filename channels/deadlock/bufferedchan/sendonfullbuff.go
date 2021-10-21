package main

import (
	"fmt"
)

func main() {
	// channel with capacity of 2
	message := make(chan string, 2)
	//problem
	// message <- "ping 1"
	// message <- "ping 2"
	// /An attempt to write to a full channel will block until
	// some other goroutine reads from it so there are no other goroutines.
	//  So when you write to the full channel, the main goroutine blocks,
	//   and since there are no other goroutines,
	//   there is no chance that the main goroutine ever can progress so its a deadlock condition.
	// message <- "ping 3"
	//solution
	//	blocked whenever the buffer is full
	go func() {
		message <- "ping 1"
		message <- "ping 2"
		message <- "ping 3" //how it works thus overflows the extra values
		// similiar to pythagoras cup https://www.youtube.com/watch?v=ISfIT3B4y6E
		message <- "ping 4"
		message <- "ping 5"
	}()
	fmt.Println(<-message)
	fmt.Println(<-message)
	fmt.Println(<-message)
	//fmt.Println(<-message)
	//fmt.Println(<-message)
}
