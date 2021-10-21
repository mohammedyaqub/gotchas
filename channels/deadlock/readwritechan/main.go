package main

import (
	"fmt"
//	"time"
)

func main() {
	out := make(chan int)
	in := make(chan int)
	//writing from one channel and reading from another channel  
	go m(in, out)
	go m(in, out)
	go m(in, out)
	go func() {
		in <- 2
		in <- 3
		in <- 4
		//defer close(out)
	}()
for v:=range out {
	//issue will wait for next value to be received but actually none 
	//there should be indication to for not waiting for more values to be received   
	fmt.Println(v)
}
	// Now we wait for each result to come in works fine but what we dont know how many value will received
	// fmt.Println(<-out)
	// fmt.Println(<-out)
	// fmt.Println(<-out)
}

func m(in <-chan int, out chan<- int) {
	fmt.Println("Initializing goroutine...")
	for {
		num := <-in
		result := num * 2
		out <- result
	}
	//close(out)
}
