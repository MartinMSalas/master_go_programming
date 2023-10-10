package main

import "fmt"

func main() {
	people := [5]string{"Ana", "Juan", "Pedro", "Maria", "Luis"}
	friends := [2]string{"Juan", "Pedro"}
	// The label is placed before the for loop
	// The label is used in the break statement
	// The break statement will break the for loop
	// and the label will break the outer for loop
outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Println("Found a friend", name, "at index", index, "!")
				// break the outer for loop
				break outer
			}
		}
	}

	fmt.Println("Next instruction after the break statement")
	i := 0
loop:
	if i<5 {
		fmt.Println("i is less than 5")
		i++
		// break the loop label
		goto loop
	}
	
/* todo:
	fmt.Println("something here") */

}
