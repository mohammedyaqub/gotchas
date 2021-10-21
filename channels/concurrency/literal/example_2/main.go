package main

import (
	"fmt"
	"time"
)

func main() {
	//simple go routine with an issue
	//	wg := sync.WaitGroup{}
	//printing most of the time last  value same  why so ?
	for i := 0; i <= 9; i++ {
		//		wg.Add(1)//the go runtime scheduler try to optimize by avoiding context swicthing between go routines
		go func() { //which might result in the prevoius holding value
			//			defer wg.Done()
			//i is declared outside the scope both the child go routine and main go routine
			// are referencing to the same var i which might be changed over time
			fmt.Println(i)//warning loop variable captured by fun literals
		}()
	}
	time.Sleep(2 * time.Second)
	//}(&wg)
	//wg.Wait()

}
