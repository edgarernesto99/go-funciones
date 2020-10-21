package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("a: ")
	fmt.Scan(&a)
	fmt.Print("b: ")
	fmt.Scan(&b)

	intercambiar(&a, &b)
	fmt.Println("a: ", a, "b: ", b)
}

func intercambiar(a, b *int) {
	aux := *a
	*a = *b
	*b = aux
}
