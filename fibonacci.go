package main

import "fmt"

func main() {
	fmt.Println(fibonacci(5))
	fmt.Println(fibonacci(6))
	fmt.Println(fibonacci(7))
	fmt.Println(fibonacci(8))
}

func fibonacci(n uint) uint {
	if n == 1 || n == 2 { //F(1) = 1
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2);
}