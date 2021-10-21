package main

import "fmt"

func main() {
	//https://stackoverflow.com/questions/36213035/why-does-go-use-rather-than-for-unary-bitwise-not
	var d uint8 = 2
	fmt.Printf("%08b\n", d)
	fmt.Printf("%08b\n", ^d) //same as ~d

	var x uint8 = 3 //0011
	var y uint8 = 4 //0100
	z := x ^ y      //0111
	fmt.Printf("%08b\n", z)
	fmt.Printf("%08b\n", ^z) //
}
