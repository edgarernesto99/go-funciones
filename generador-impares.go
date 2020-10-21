package main

import "fmt"

func main() {
	nextodd := generadorImpares()
	fmt.Println(nextodd())
	fmt.Println(nextodd())
	fmt.Println(nextodd())
	fmt.Println(nextodd())
}

func generadorImpares() func() uint {
	i := uint(1)
	return func() uint {
		var odd = i
		i += 2
		return odd
	}
}