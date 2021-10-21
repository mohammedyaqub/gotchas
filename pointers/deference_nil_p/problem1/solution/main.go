package main

import (
	"fmt"
)

func main() {
	var pointer *int
	if pointer != nil {
		fmt.Println(*pointer)
	} else {
		//allocate a memory to hold value
		pointer = new(int)
	}
	*pointer = 10
	fmt.Println(*pointer)
}
