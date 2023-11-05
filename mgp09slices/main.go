package main

import "fmt"

func main() {
	
	var number [5]int;

	fmt.Printf("%v\n", number)
	fmt.Printf("%#v\n", number)

	var a1 = [3]int{1, 2, 3}

	fmt.Printf("%v\n", a1)
	fmt.Printf("%#v\n", a1)


	a3 := [...]int{1, 2, 3}
	fmt.Printf("%v\n", a3)

	fmt.Printf("%#v\n", a3)
	fmt.Printf("Length: %d\n", len(a3))

	a5 := [...]float64{67.7, 89.8, 21, 78,}
	fmt.Printf("%#v\n", a5)
	fmt.Println(a5)

	for i, v :=  range a5 {
		fmt.Printf("%d the element of a5 is %.2f\n", i, v)
	}

	for i := 0; i < len(a5); i++ {
		fmt.Printf("%d the element of a5 is %.2f\n", i, a5[i])
	}

	balances := [2][3]int{
		{5,6,7},
		{8,9,10},
	}
	fmt.Println(balances)

	m := [3]int{1, 2, 3}
	n := m
	fmt.Println(m, n)
	n[0] = 100
	fmt.Println(m, n)
	fmt.Printf("m: %p\n", &m)

	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,

	}
	fmt.Println(grades)

	accounts := [3]int{2: 50}
	fmt.Println(accounts)

	names := [...]int{
		5: 1,
	}
	fmt.Println(names, len(names))
	cities := [...]string{
		5: "Paris",
		"Catamarca",
		"New York",
		1: "London",
	}
	for i,v := range cities {
		fmt.Printf("%d the element of a5 is %s\n", i, v)
	}	
	fmt.Println(cities, len(cities))
}