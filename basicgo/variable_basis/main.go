package main

import "fmt"

const secondsInHour = 3600

func main() {
	var age int = 30
	fmt.Println("Age is", age)
	var name= "World"
	fmt.Println("Hello",name)
	

	fmt.Print("The seconds in one Hour is: ", secondsInHour, "\n")
	s := "fun!"
	fmt.Println("Learning Go is", s, "Go is", s, "Go is", s)
	name,trait := "Pocholo","King of the Night"
	fmt.Println("Hello",name,"you are the",trait)
  car,cost := "Audi", 50000
  fmt.Println(car, "is", cost)
	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary float64
		name1 string
		age1 int
		gender bool		
	)
	fmt.Println(salary,name1,age1,gender)

	var i,j int
	i,j = 5,8
	fmt.Println(i,j)
	i,j = j,i
	fmt.Println(i,j)
	sum := 5 + 2.3
	fmt.Println(sum)
}
