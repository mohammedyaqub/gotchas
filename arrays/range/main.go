package main
import "fmt"


func main(){

	s:=[]int{1,2,3,4}
	for i,v:=range s{
		fmt.Println(i ," ->",v)
	}

	// str:=[]string{"a","b","c"}
	// for val:=range str{//wrong way writing this because first value is index and second value as values
	// 	fmt.Println(val)
	// }
	
	str:=[]string{"a","b","c"}
	for _,val:=range str{//correct first value we are not caring about that 
		fmt.Println(val)
	}
}