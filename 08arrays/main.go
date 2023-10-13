package main

import "fmt"

func main() {

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{5, -3, 5}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"Dan", "Diana", "Paul", "John"}
	fmt.Printf("%#v\n", a3)

	a4 := [4]string{"x", "", "y"}
	fmt.Printf("%#v\n", a4)

	a5 := [...]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%#v\n", a5)
	fmt.Printf("a5 length is %d\n", len(a5))

	numbers := [8]int{3, 16, -2, 10, 23, 12, 8, 6}
	fmt.Printf("%#v\n", numbers)

	for i := 0; i < len(numbers); i++ {
		fmt.Printf("numbers[%d] is %d\n", i, numbers[i])
	}
	for i, v := range numbers {
		fmt.Printf("numbers[%d] is %d\n", i, v)
	}

	balances := [2][3]int{
		{5, 6, 7},
		[3]int{8, 9, 10},
	}
	fmt.Printf("%#v\n", balances)

	grades := [3]int{
		1: 10,
		0: 5,
		2: 7,
	}
	fmt.Printf("%#v\n", grades)
	weekend := [7]bool{5: true, 6: true}
	fmt.Printf("%#v\n", weekend)

	
}
