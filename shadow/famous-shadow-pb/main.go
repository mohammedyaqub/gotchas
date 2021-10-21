package main

import (
	"fmt"
	"strconv"
)



func main() {
	//s = "amer"
	var n int
	s := "124"
	if n, err := strconv.Atoi(s); err != nil {//here is the problem n is created again with scope of if block 
		fmt.Println("not converted", n)
	}
	fmt.Println(s, n)//124 0 not expected both should be same 
}
