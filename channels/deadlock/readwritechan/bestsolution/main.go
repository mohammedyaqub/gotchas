package main

import (
	"fmt"
	"time"
)
// You don’t want leftover channels and goroutines to consume resources and cause
// leaky applications. You want to safely close channels and exit goroutines.
func main() {
	out := make(chan int)
	in := make(chan int)
	done :=make(chan bool)
	until := time.After(5 * time.Millisecond)//will create and initialize the channel and return 
	//writing from one channel and reading from another channel  
	go m(in, out,done)
	go m(in, out,done)
	go m(in, out,done)
	go func() {
		in <- 2
		in <- 3
		in <- 4
		//defer close(out)
	}()
// 	This gives both sides the opportunity to cleanly handle
// the closing of the channel
for {
	select {
	case v:=<-out:
		fmt.Println(v)
	case <-until:
		fmt.Println("we are done with task so please give indicate to sender ")
		done <- true
		close(out)
		return
}
}

}
//If the receiver hits a stopping condition, it must let the sender know
func m(in chan int, out chan<- int,done chan bool) {
	fmt.Println("Initializing goroutine...")
	for {
		select {
		case num := <-in:
		result := num * 2
		out <- result
		case <-done:
			close(in)
			fmt.Println("both channels are close succesfully and will be gc")
	}
}

}