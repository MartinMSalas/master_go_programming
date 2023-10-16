package main

import "fmt"

func main() {
	//var cities []string
	//fmt.Printf("%#v\n", cities)

	cities := []string{"Paris", "London", "Tokyo", "Catamarca", "Cordoba"}

	fmt.Printf("%#v\n", cities)
	cities[4] = "New York"
	fmt.Printf("%#v\n", cities)

	for index, value := range cities {
		fmt.Printf("cities: %d is %s\n", index, value)
	}

	var n []int
	fmt.Println(n == nil)

	m := []int{}
	fmt.Println(m == nil)
	fmt.Println(m)
	a, b := []int{1, 2, 3}, []int{1, 2, 3, 4}
	fmt.Println(len(a) + len(b))

	var eq bool = true

	if len(a) != len(b) {
		eq = false
	} else {

		for i := range a {
			if a[i] != b[i] {
				eq = false
				break
			}
		}
	}
	if eq {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}

}
