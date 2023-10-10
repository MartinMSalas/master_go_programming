package main // package main

import "fmt"

//const secondsInHour = 3600 // const secondsInHour = 3600

func main() { // func main()
	fmt.Println("Hello World")                                           // fmt.Println("Hello World")
	distance := 60.8                                                     // distance in km
	fmt.Println("The distance in km is", distance)                       // fmt.Printf("The distance in km is %v\n", distance)
	fmt.Printf("The distance in miles is %f\n", float64(distance)*0.621) // fmt.Printf("The distance in miles is %v\n", float64(distance)*0.621)
}
