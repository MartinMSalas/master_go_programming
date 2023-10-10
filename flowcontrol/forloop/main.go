package main

import "fmt"

func main() {
	acum := 0
	for i := 0; true; i++ {
		if i%13 == 0 {
			fmt.Println(i)
			acum++
			if acum == 13 {
				break
			}
		}
	}
}
