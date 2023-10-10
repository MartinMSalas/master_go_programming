package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {

	//a := int8(128)
	//fmt.Println(a)

	var b int8 = 127
	fmt.Printf("%d\n", b+1)

	f := float32(math.MaxFloat32)
	fmt.Println(f)
	f = f * 1.2
	fmt.Println(f)

	x := 3
	y := 4.5

	fmt.Println(float64(x) * y)
	x = x * int(y)
	fmt.Println(x)

	s := string(99)
	fmt.Println(s)

	var myStr = fmt.Sprintf("%f", 3.14)
	fmt.Println(myStr)
	myStr = fmt.Sprintf("%d", 100)
	fmt.Println(myStr)

	s1 := "3.123"
	fmt.Printf("%T\n", s1)
	fmt.Println(s1)
	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1)
	fmt.Printf("%T\n", f1)

	i,err := strconv.Atoi("-50")
	_ = err
	s2 := strconv.Itoa(-50)
	fmt.Println(i)
	fmt.Printf("Type: %T and value %v\n", i,i) 
	fmt.Printf("Type: %T and value %v\n", s2,s2)

}
