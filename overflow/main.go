package main

import (
	"fmt"
	"math"
)

func main() {
	var x uint8 = 255
	x++ // 0
	fmt.Println(x)

	//a := int8(128)
	//fmt.Println(a)

	var b int8 = 127
	fmt.Printf("%d\n", b+1)

	f := float32(math.MaxFloat32)
	fmt.Println(f)
	f = f * 1.2
	fmt.Println(f)
}
