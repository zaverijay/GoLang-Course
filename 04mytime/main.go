package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")

	presentTime := time.Now()
	fmt.Println("Current time is : ", presentTime.Format("01-02-2006"))

	createdDate := time.Date(2001, 7, 12, 10, 55, 50, 00, time.Local)
	fmt.Println("Created date is ,", createdDate.Format("Monday"))
}
 