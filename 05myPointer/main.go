package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointer module")

	myNumber := 46

	var ptr = &myNumber

	fmt.Println("the value of ptr is, ", ptr)
	fmt.Println("the value of ptr is, ", *ptr)
	
	*ptr = *ptr + 2
	
	fmt.Println("the value of ptr is, ", myNumber)

}
