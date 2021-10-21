package main
import ("fmt")



func main(){
	 var pointer *int
	 fmt.Println(pointer==nil)
	 *pointer=2// invoked func fatalpanic nil pointer dereference"
	 
	 fmt.Println(*pointer)
}