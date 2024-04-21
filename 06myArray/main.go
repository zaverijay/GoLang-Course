package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array module")

	var goSubjects [4]string

	goSubjects[0] = "GOPATH"
	goSubjects[1] = "GOROOT"
	goSubjects[2] = "GOOS"
	goSubjects[3] = "GOGC"

	fmt.Println("the main subjects in GO are :", goSubjects)

	var typeList = [4]string{"integer", "string", "float", "boolean"}
	fmt.Println("the main types in GO are :", typeList)

	
}
