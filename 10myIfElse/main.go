package main

import "fmt"

func main() {
	fmt.Println("Welcome to IF-Else module")

	myNumber := 20
	var result string

	if myNumber > 10 {
		result = "the number is greater than 10"
	} else if myNumber < 10 {
		result = "the number is less than 10"
	} else {
		result = "the number is 10"
	}

	fmt.Println(result)

	if 9%2 == 0 {
		fmt.Println("the number is even")
	} else { 
		fmt.Println("the number is odd")
	}

	if n := 11; n <= 10 {
		fmt.Println("the n is less than or equal to 10")
	} else {
		fmt.Println("the number is greater  than 10")
	}
}
