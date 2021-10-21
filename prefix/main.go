package main

import "fmt"
//Unlike other languages,
 //Go doesn't support the prefix version of the operations. 
 //You also can't use these two operators in expressions.
func main() {  
    data := []int{1,2,3}
    i := 2
    i--
    fmt.Println(data[i])
    i++
    fmt.Println(data[i])
	i1 := 2
    i1--
    fmt.Println(data[i1])
	i2 := 2
    //--i2 syntax error: unexpected --, expecting }
    fmt.Println(data[i2])
}