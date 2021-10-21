package main

import (
	"fmt"
	"sync"
)
func main() {
	var wg sync.WaitGroup
for _, salutation := range []string{"hello", "greetings", "good day"} {//order is not important if so then use unbuffered chan
 wg.Add(1)
 go func(salutation string) {
 defer wg.Done()
 fmt.Println(salutation)
 }(salutation)
}
wg.Wait()
}
