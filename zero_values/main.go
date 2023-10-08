package main

import "fmt"

//const secondsInHour = 3600
 
func main() {
	const (
		min1 = 500
		min2
		min3 
	)
	fmt.Println("Minutes", min1, min2, min3)
	a, b := 4, 5.2
	a = int(b)

	fmt.Println("a is", a, "b is", b)
}
