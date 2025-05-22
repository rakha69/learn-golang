package main

import "fmt"

func main() {
	var names [4]string

	names[0] = "Hafidz"
	names[1] = "Muhammad"
	names[2] = "Rakha"
	names[3] = "Shidqi"
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	fmt.Println(names[3])

	var values = [...]int{
		75,
		80,
		85,
	}
	fmt.Println(values)
	fmt.Println(len(values))
}
