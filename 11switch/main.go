package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch module")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(7)

	fmt.Println("the dice number is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("the number is one")
	
	case 2:
		fmt.Println("the number is two")
	
	case 3:
		fmt.Println("the number is three")
	
	case 4:
		fmt.Println("the number is four")
	
	case 5:
		fmt.Println("the number is five")
	
	case 6:
		fmt.Println("the number is six")
	}
	
}

