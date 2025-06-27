package main

import "fmt"

func getCompleteName() (firstname, middlename, lastname string) {
	firstname = "Hafidz"
	middlename = "Muhammad"
	lastname = "Rakha"

	return firstname, middlename, lastname
}

func main() {
	a, b, c := getCompleteName()
	fmt.Println(a, b, c)
}
