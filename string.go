package main

import (
	"fmt"
	"strings"
)

func main() {
	str := " i love my country"
	fmt.Println("length of given string", len(str))
	fmt.Println("Ascii value of first alp", "A"[0])
	fmt.Println(strings.ToUpper(str))
	fmt.Println(strings.ToLower(str))
	fmt.Println(strings.HasPrefix(str, "IN"))
	fmt.Println(strings.HasSuffix(str, "ry"))
	fmt.Printf(strings.Repeat(str, 2))

}
