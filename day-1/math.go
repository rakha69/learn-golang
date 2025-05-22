package main

import "fmt"

func main() {
	var a = 100
	var b = 20
	var c = a / b
	fmt.Println(c)

	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	i += 5 // i = i + 5
	fmt.Println(i)

	var j = 5
	j++ // j = j + 1
	j++ // j = j + 1
	fmt.Println(j)
}
