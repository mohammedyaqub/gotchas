package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	done := make(chan int) //used for signallng the go routines we are done so please deallocate the resources
	var v2 int
	go func() {
		v := 1
		ch1 <- v
		v2 := <-ch2
		fmt.Println(v, v2)
		//close(ch1)
		done <- 0
	}()
	v := 2
	//var v2 int
	// When using a for-select loop, you must include a way to exit the loop.
	_ = v2
	for {
		select {

		case ch2 <- v:
		case v2 = <-ch1:
			//	case c<-true :ls
		case <-done:
			return
		}
		//	close(ch2)
	}
	//_=v2
	//	<-c
	//close(done)
	//fmt.Println(v, v2)
}
