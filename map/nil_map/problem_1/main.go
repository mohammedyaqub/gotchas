package main

import (
	"fmt"
)

func main() {
	//correct way of doing
	// m := make(map[string]int)
	// m["start"] = 100
	// //fmt.Println(m["start"])
	// idx, key := m["start"]
	// fmt.Println(idx)
	// fmt.Println(key)
	//wrong way
	var m map[string]int
	m["start"]=100
	fmt.Println(m["start"])

}
