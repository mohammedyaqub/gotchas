package main
import ("fmt")


func main() {
	var s  interface{}='a'//ascii 97
    if res, ok := s.(int); ok {
        fmt.Println("expecting integer type value =>",res)
    } else {
        fmt.Println("whatever the interface holding type value =>",s) 
    }
}