package main

//https://go.dev/blog/defer-panic-and-recover
import "fmt"

func c() (i int) {
	defer func() { i = i + 10 }() //evalute but execute just before return statement here no furthere statement to execute
	//i=i+30 reduntant bcz later time  i is going to modify again
	return 20
}
func b() {
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}
func a() {
	i := 0
	defer fmt.Println(i)
	i = i + 2
	defer fmt.Println(i)
	// return
}
func main() {
	var i int = 1
	//A deferred function’s arguments are evaluated when the defer statement is evaluated.
	//Each time the "defer" statement executes, the function value and parameters to the call
	//are evaluated as usual and saved a new but the actual function body is not executed.
	//means whenever there is defer keyword at that time only the function parameters are evaluted but execution happens later
	defer fmt.Println("result =>", func() int { return i * 2 }()) //result => 2
	a()
	b()
	fmt.Printf("\n%v", c())
	i = i + 10
	fmt.Printf("\n%v\n", i) //11
}

/*
output
2
0
3210
30
11
result => 2
 best uses of defer deallocate or closed the resources irrespective of the returning if something went wrong*/
