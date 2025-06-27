package main

import "fmt"

func sayHelloTo(firstname string, lastname string) {
	fmt.Println("hello", firstname, lastname)
}

func main() {
	sayHelloTo("Hafidz", "muhammad")
}
