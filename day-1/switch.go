package main

import "fmt"

func main() {
	name := "Rakha"

	switch name {
	case "Rakha":
		fmt.Println("Hello Rakha")
	case "Kiwil":
		fmt.Println("Hello Kiwil")
	default:
		fmt.Println("Hai, boleh kenalan?")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	name = "muhammad"
	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
