package main

import "fmt"


func main() {
	const (
		c1 = iota
		c2 
		c3 
	)

	fmt.Println(c1, c2, c3)


	const d = 5 > 10 // untyped boolean constant
	fmt.Println(d)
	/*
	const x = 5
	const y = 2.2 * x // x is converted to float64
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	var z float64 = x * 1.2 // x is converted to float64
	fmt.Println(z)
*/
	const (
		North = iota // 0
		East // 1
		South // 2
		West // 3
	)

	fmt.Println(North, East, South, West)

	const (
		x = -(iota + 2) // -2
		_ // -3
		y // -4
		z // -5

	)
	fmt.Println(x, y, z)

}