package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to Maps module")

	languages := make(map[string]string)

	languages["JS"] = "JavaScript"
	languages["RB"] = "Ruby"
	languages["PY"] = "Python"

	fmt.Println("Languages are ", languages["JS"])

	// delete(languages, "RB")
	fmt.Println("Languages are ", languages)

	for key, value := range languages {
		fmt.Printf("for key %v, value is %v\n", key, value)
	}

}
