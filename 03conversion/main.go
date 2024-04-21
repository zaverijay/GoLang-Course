package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to conversion module")

	fmt.Println("Please rate out facility:")

	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	
	fmt.Println("Thanks for rating, ", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)

	if err != nil {
		fmt.Println("Error is: ", err)

	} else {
		fmt.Println("We have added 1 to your rating", numRating + 1)

	}

}
