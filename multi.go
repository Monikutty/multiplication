package main

import "fmt"

func main() {

	var a = [3][3]int{{10, 20, 30}, {40, 50, 60}, {70, 80, 90}}
	var i, j int

	for i = 0; i < 3; i++ {
		for j = 0; j < 3; j++ {
			fmt.Print(a[i][j], "\t")
		}
		fmt.Println()
	}
}
