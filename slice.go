package main

import "fmt"

func main() {
	slice1 := []int{2, 4, 6, 8, 10, 12, 14}
	printSlice(slice1)

	slice2 := slice1[:0]
	printSlice(slice2)

	slice3 := slice1[:4]
	printSlice(slice3)

	slice4 := slice1[2:]
	printSlice(slice4)
}
func printSlice(s []int) {
	fmt.Printf("length =%d %v\n", len(s), s)
}
