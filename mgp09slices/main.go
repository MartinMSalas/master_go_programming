package main

import "fmt"

func main() {
	//var cities []string
	//fmt.Printf("%#v\n", cities)

	cities := []string{"Paris", "London", "Tokyo", "Catamarca", "Cordoba"}

	fmt.Printf("%#v\n", cities)
	cities[4] = "New York"
	fmt.Printf("%#v\n", cities)

	for index,value := range cities {
		fmt.Printf("cities: %d is %s\n", index, value)
	}
}
