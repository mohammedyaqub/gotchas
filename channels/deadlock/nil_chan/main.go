package main

import (  
    "fmt"
)
//Send and receive operations on a nil channel block forever.
func main() {
//declare a nil channel
//var ch chan int// initialize to nil
ch := make(chan int)//works fine
go func (ch chan int){
ch <- 23
}(ch)
fmt.Println(<-ch)

}