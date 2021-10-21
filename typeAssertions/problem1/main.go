package main

import "fmt" 
//Failed type assertions return the "zero value" for the target
//type used in the assertion statement. This can lead to
//unexpected behavior when it's mixed with variable shadowing.
func main() {  
   // var s  interface{}=8//works
	var s  interface{}="gopher"
    if s, ok := s.(int); ok {//variable   new s declared inside if block
        fmt.Println("expecting integer type =>",s)
    } else {
        fmt.Println("whatever the interface holding type  =>",s) 
    }
}