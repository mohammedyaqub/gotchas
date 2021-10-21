package main

import (
	"fmt"
)

//arrays in go passed as value not pointer
//if you need to update pass as reference
//reflect the original array
func UpdateArray(arr *[3]int) *[3]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] + 2
	}
	//fmt.Println(arr)
	return arr
}
//doesn't reflect the original array data
func Update(arr [3]int) [3]int {
	for i := 0; i < len(arr); i++ {
		arr[i] = arr[i] + 2
	}
	return arr
}
func main() {
	arr := [3]int{1, 2, 3}
	fmt.Println(arr) //1 2 3
	Update(arr)
	fmt.Println(arr)//still 1 2 3
	retArr := UpdateArray(&arr)
	fmt.Println(retArr) //3,4,5
	fmt.Println(arr)    //3 4 5
	// func(arr [3]int) {
	// 	arr[0] = 23
	// 	fmt.Println(arr)//23 2 3
	// }(arr)
	// fmt.Println(arr) 1 2 3
}
