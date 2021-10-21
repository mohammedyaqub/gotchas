package main

import "fmt"

func main() {
	//fmt.Println(~3) //illegal operator in go not allowed it can be achieved by ^ xor works for both unary and binary operator
	var d uint8 = 4
	fmt.Printf("%08b\n", d)
	fmt.Printf("%08b\n", ^d)
	//output 
	//00000100
	//11111011
	fmt.Printf("\n%d",^d)// 255 - 4 = 251
}
