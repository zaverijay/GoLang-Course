package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops module")

	days := []string{"Monday", "Tuesday", "Webneday", "Thursday"}

	fmt.Println("the days are", days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index, value := range days {
	// 	fmt.Printf("the index is %v and day is %v\n", index, value)
	// }

	roughValue := 1

	for roughValue <= 10 {

		// if roughValue == 5 {
		// 	fmt.Println("your loop is going to break")
		// 	roughValue++
		// 	break
		// }

		// if roughValue == 5 {
		// 	// fmt.Println("your loop is going to skip 5")
		// 	roughValue++
		// 	continue
		// }

		if roughValue == 5 {
			// fmt.Println("your loop is going to skip 5")
			roughValue++
			goto newLoc
		}
		fmt.Println("the value is ", roughValue)
		roughValue++
	}

newLoc:
	fmt.Println("jumping here")
}
