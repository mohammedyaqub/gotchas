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

func main() {
	p := person{
		FirstName:  "Pat",
		MiddleName: "Perry", // This line won't compile
		LastName:   "Peterson",
	}
	fmt.Println(p.MiddleName)
}
