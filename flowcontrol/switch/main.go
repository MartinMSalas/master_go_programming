package main

import "fmt"

func main() {
	language := "golang"

	switch language {
	case "Python":
		println("You are learning Python! You are cool!")
	case "Go", "golang":
		println("You are learning Go! That's awesome!")
	default:
		println("Any other programming language is a good start!")
	}

	n := 4
	switch true {
	case n%2 == 0:
		println("Even")
		fmt.Println("the value of n is", n)
	case n%2 != 0:
		println("Odd")

	}
}
