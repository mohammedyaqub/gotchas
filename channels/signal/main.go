package main

import (
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	//c := make(chan os.Signal)// issue := the channel used with signal.Notify should be buffered
	c := make(chan os.Signal,1)//no issue
	signal.Notify(c, os.Interrupt)

	time.Sleep(2 * time.Second)

	// Block until a signal is received.
	s := <-c
	fmt.Println("Got signal:", s)
}