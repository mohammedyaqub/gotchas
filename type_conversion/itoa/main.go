package main

import "fmt"

func main() {
//wrong way 
	//The string supports type conversion from int, here string() will treat integer as rune. The rune of 97 is a.
	a := 97
	s := string(a)
	fmt.Println(s)
}
