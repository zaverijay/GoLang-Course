package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to Array slice module")

	var goTypes = []string{"intger"}
	// fmt.Printf("The type of goTypes is %T\n", goTypes)
	
	goTypes = append(goTypes, "string", "float", "boolean")
	
	fmt.Println("the goTypes are", goTypes)

	highScores := make([]int, 4)

	highScores[0] = 699
	highScores[1] = 199
	highScores[2] = 99
	highScores[3] = 799

	highScores = append(highScores, 899, 599, 1099)

	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println(highScores)

	// using the slieces to remove th element by index
	var index = 2
	highScores = append(highScores[:index], highScores[index+1:]...)
	fmt.Println(highScores)
}
