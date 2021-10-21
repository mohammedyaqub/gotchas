package main

import (
"fmt"
//""
)

type person struct {
	FirstName  string
	MiddleName *string
	LastName   string
}

func stringp(s string) *string {
	return &s
}
//Use a helper function to turn a constant value into a pointer
func main() {
	p := person{
		FirstName:  "Pat",
		MiddleName: stringp("Perry"), // This line won't compile
		LastName:   "Peterson",
	}
	fmt.Println(*p.MiddleName)
}
