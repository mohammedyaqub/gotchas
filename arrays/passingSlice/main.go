package main

import (
	"fmt"
)
/*
https://stackoverflow.com/questions/39993688/are-slices-passed-by-valuetype 
SliceHeader struct {
    Data uintptr
    Len  int
    Cap  int
}*/

//slice are passed by value too but slice includes slice 
//header which contain pointer 
//  to the same memory location of backing array
func Update(s []int) []int {
	for i := 0; i < len(s); i++ {
		s[i] = s[i] + 2
	}
	return s
}
func main() {
	s1 := []int{1, 2, 3}
	fmt.Println(s1)
	retslice := Update(s1)
	fmt.Println(retslice)
	fmt.Println(s1)
}
