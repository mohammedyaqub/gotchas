package main

import (
	"fmt"
	"strconv"
)

var (
	n   int
	err error
)

func main() {
	//s = "amer"
	//var n int
	
	s := "124"
	if n, err = strconv.Atoi(s); err != nil {
		fmt.Println("not converted", n)
	}
	fmt.Println(s, n)
}
