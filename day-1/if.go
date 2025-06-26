package main

import "fmt"

func main() {
	name := "Rakha"

	if name == "Rakha" {
		fmt.Println("Hallo Rakha")
	} else if name == "Hafidz" {
		fmt.Println("Hallo Hafidz")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	length := len(name)
	if length > 5 {
		fmt.Println("Nama Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}
}
