package main

import "fmt"

//import "os"
var a, b int //global var
//returning multiple values from a chan of func
func f(c chan func() (int, int, int)) {
	//var a, b int
	fmt.Scanln(&a, &b)//2 2
	c <- (func(i int, s int) func() (int, int, int) {//writing the results to the channels
		return func() (int, int, int) {
			return i + s, i - s, i * s
		}
	}(a, b))

}

func main() {
	c := make(chan func() (int, int, int))
	go f(c)
	add, sub, mul := (<-c)()
	fmt.Printf("%v + %v = %v\n", a, b, add)
	fmt.Printf("%v - %v = %v\n", a, b, sub)
	fmt.Printf("%v * %v = %v\n", a, b, mul)
	// result
	// 2 2 input
	//output 
	// 2 + 2 = 4
	// 2 - 2 = 0
	// 2 * 2 = 4
}
