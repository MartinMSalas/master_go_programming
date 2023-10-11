package main

import "fmt"

const scopeLevel = "package"

func main() {
	fmt.Println("scopeLevel:", scopeLevel)
	
	const scopeLevel = "function"
	fmt.Println("scopeLevel:", scopeLevel)
	f1()
}

func f1(){
	fmt.Println("scopeLevel:", scopeLevel)
}