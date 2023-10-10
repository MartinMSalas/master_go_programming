package main

import "fmt"

func main() {
	a, b := 1, 8
	fmt.Println("a is", a, "b is", b)

	a += b
	fmt.Println("a is", a, "b is", b)
	b -= 2
	fmt.Println("a is", a, "b is", b)
	a *= 3
	fmt.Println("a is", a, "b is", b)
	b /= 2
	fmt.Println("a is", a, "b is", b)
	
}
