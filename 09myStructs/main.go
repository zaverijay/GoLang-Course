package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs module")

	jay := User{"jay", "jay@go.dev", true, 23}
	fmt.Println("the user name is ", jay.Name)
	fmt.Printf("the user details are %+v\n", jay)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
