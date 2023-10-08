package main

import "fmt"

func main() {
	// INT Type

	var i1 int8 = 127

	fmt.Printf("%T\n", i1)

	var i2 uint64 = 99999
	fmt.Printf("%T\n", i2)
	// FLOAT Type

	var f1, f2, f3 float32 = 3.14, -2.15, 1.0
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	// RYNE BYTE Type

	var b1 byte = 'f'
	var r1 rune = 1351
	fmt.Printf("%T %T\n", b1, r1)
	fmt.Print(b1, " ", r1, "\n")

	// Array Type
	var numbers = [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", numbers)
	// Slice Type
	var cities = []string{"London", "Tokyo", "New York"}
	fmt.Printf("%T\n", cities)

	// Map Type
	balances := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,
		"CHF": 666.11,
	}
	fmt.Printf("%T\n", balances)

	// Struct Type
	type Person struct {
		name string
		age  int
	}
	var you Person
	fmt.Printf("%T\n", you)

	// Pointer Type
	var x int = 2
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v and address of %v\n", ptr, ptr, &ptr)

	// Function Type
	fmt.Printf("%T\n", f)
}

func f() {}
