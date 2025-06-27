package main

import "fmt"

func getFullName() (string, string) {
	return "rakha", "qila"
}

func main() {
	//firstname, Lastname := getFullName()
	//fmt.Println(firstname, lastname)

	firstname, _ := getFullName()
	fmt.Println(firstname)
}
