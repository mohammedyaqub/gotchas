package main

import (
	"fmt"
	"strconv"
)

func main() {
	integer := 97 
	str:= strconv.Itoa(integer)
	fmt.Println(str)
	
	s := fmt.Sprintf("%d", integer)
	fmt.Println(s)
}