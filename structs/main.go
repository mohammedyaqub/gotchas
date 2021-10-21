package main
//setting the field value in wrong way using short declartion :=
import (  
  "fmt"
)
//devs.cloudimmunity.com/gotchas-and-common-mistakes-in-go-golang/index.html#opening_braces
type info struct {  
  result int
}

func work() (int,error) {  
    return 13,nil  
  }

func main() {  
  var data info
var err error
//   data.result, err := work() //error
//   fmt.Printf("info: %+v\n",data)
//correct way 
data.result, err = work() 
if err != nil {
fmt.Println(err)
}
var s []int
    s = append(s,1)
fmt.Println(s)
 }