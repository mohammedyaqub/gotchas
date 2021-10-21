package main

import (
  "fmt"
)

const (
  zero = iota
  one  = iota
)
//The First Use of iota Doesn't Always Start with Zero
const (
  
  info  = "processing"//0
  data ="engg"//1
  math ="interesting"//2
  bZero = iota//3 not start with 0
  bOne  = iota//4
)
//iota is really an index operator for the current line in the constant declaration
 //block, so if the first use of iota is not the first line in the constant declaration
  //block the initial value will not be zero.
func main() {
  fmt.Println(zero,one) //prints: 0 1
  fmt.Println(bZero,bOne) //prints: 2 3
}