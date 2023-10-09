package main

import "fmt"

type km float64
type mile float64

func main() {
	type speed uint
	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s1 + s2)
	fmt.Printf("Type of s1 is %T\n", s1)

	var parisToLondon km = 465
	var distanceInMiles mile
	distanceInMiles = mile(parisToLondon) / 0.621
	fmt.Println(distanceInMiles)
}
