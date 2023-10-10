package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println("os.Args:", os.Args)
  fmt.Println("Path:", os.Args[0])
	fmt.Println("1st argument:", os.Args[1])

	fmt.Println("No. of items inside os.Args:", len(os.Args))

	result, _ := strconv.ParseFloat(os.Args[1], 64)
	fmt.Printf("The type of result is %T and it value %.2f \n", result,result)

	if i, err := strconv.ParseFloat(os.Args[1], 64); err == nil {
		fmt.Printf("The type of i is %T and it value %.2f\n", i,i)
	}else{
		fmt.Println(err)
	}

	if args := os.Args; len(args) != 2 {
		fmt.Println("One argument is required!")
	}else if km,err := strconv.Atoi(args[1]); err != nil {
		fmt.Println("The argument must be an integer (e.g. 10)")
	}else {
		fmt.Printf("%d km in miles is %.2f\n", km, float64(km)*0.621)
	}
}