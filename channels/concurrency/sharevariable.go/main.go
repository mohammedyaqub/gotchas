package main

import (
	"fmt"
	"sync"
)
func main() {
	var wg sync.WaitGroup//1
salutation := "hello"//2
wg.Add(1)//3
go func() {// 4this  thread runs independently regardless of main thread
 defer wg.Done()
 salutation = "welcome"
}()
fmt.Println(salutation)// 5 hello
wg.Wait()//6 wait to join till above thread finished 
fmt.Println(salutation)//7 welcome
}
