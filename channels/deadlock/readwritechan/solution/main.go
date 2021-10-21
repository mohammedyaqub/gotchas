package main

import (
	"fmt"
	"time"
)

func main() {
	out := make(chan int)
	in := make(chan int)
	until := time.After(5 * time.Millisecond) //will create and initialize the channel and return
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
	for {
		select {
		case v := <-out:
			fmt.Println(v)
		case <-until:
			fmt.Println("we are done with task so please close the channel")
			close(out)
			return
		}
	}

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
