package main

//Using a for loop and an if statement print out all the numbers between 1 and 50 that divisible by 7.
import "fmt"

func main() {
	acum := 0
	for i := 1; i <= 500; i++ {
		if i%7 == 0 && i%5 == 0 {
			println(i)
			acum++
			if acum == 3 {
				break
			}
		}
	}
	for i := 1984; i <= 2023; {
		println(i)

		i++
	}
	age := 39
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
