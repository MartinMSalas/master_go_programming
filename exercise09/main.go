package main

import "fmt"

func main() {

	var number [5]int

	fmt.Printf("%v\n", number)
	fmt.Printf("%#v\n", number)

	var a1 = [3]int{1, 2, 3}

	fmt.Printf("%v\n", a1)
	fmt.Printf("%#v\n", a1)

	//	Using the var keyword, declare an array called cities with 2 elements of type string. Don't initialize the array.
	var cities [2]string

	//Print out the cities array and notice the value of its elements.
	fmt.Printf("%v\n", cities)

	//2. Declare an array called grades of type [3]float64 and initialize only the first 2 elements using an array literal.
	grades := [3]float64{1, 2}

	//Print out the grades array and notice the value of its elements.
	fmt.Printf("%v\n", grades)
	//3. Declare an array called taskDone using the ellipsis operator. The elements are of type bool. Print out taskDone.
	taskDone := [...]bool{true, false, true}
	fmt.Printf("%v\n", taskDone)

	//4. Iterate over the array called cities using the classical for loop syntax and the len function and print out element by element.
	for i := 0; i < len(cities); i++ {
		fmt.Println(cities[i])
	}
	//5. Iterate over grades using the range keyword and print out element by element.
	for i, v := range grades {
		fmt.Println(i, v)
	}
	nums := [...]int{30, -1, -6, 90, -6}

	// exercise 2
	for i,v := range nums{
		if(v>0 && v%2==0){
			fmt.Println(i,v)
		}
	}

	myArray := [4]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10.0
	myArray[0] = a

	myArray[3] = 10.10

	fmt.Println(myArray)
}
