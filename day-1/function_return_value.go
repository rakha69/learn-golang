package main

import "fmt"

func getHello(name string) string {
	hello := "Hello" + name
	return hello
}

func main() {
	result := getHello("Rakha")
	fmt.Println(result)

	fmt.Println(getHello("kiwil"))
	fmt.Println(getHello("qila"))
}
