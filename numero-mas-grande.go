package main

import "fmt"

func main() {
	s := []int{1,7,3,5,23,12,4}
	fmt.Println(max(s...))
}

func max(args... int) int {
	mayor := args[0]
	for _, v := range args {
		if v > mayor {
			mayor = v
		}
	}
	return mayor
}