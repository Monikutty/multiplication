package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
	age       int
}

func main() {
	x := person{age: 23, firstName: "monisha", lastName: "murugesan"}
	fmt.Println(x)
	fmt.Println(x.firstName)
}
