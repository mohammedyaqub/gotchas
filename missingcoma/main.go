package main

import "fmt"
//myvar := 1 //error should be declared inside {}

func main() {
//{ //error wrong way of doing , can't have the opening brace on a separate line
    fmt.Println("hello there!")//automatically insert ; done by compiler 
}
// the compiler will treat the above code as 
// func main();which is invalid in golang
//  {

// 		fmt.Println("hello there!");
// 	};

// func main() {  
//     x := []int{
// 		11,
// 		//222 //error
// 		33,//correct way
// 	 }
//     _ = x
// }