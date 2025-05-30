package main

import "fmt"

func main() {
	//cara lama
	//var person map[string]string = map[string]string
	//person["name"] = "Rakha",
	//person["address"] = "Tangerang",

	//cara cepat
	person := map[string]string{
		"name":    "Rakha",
		"address": "Tangerang",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person)

	book := make(map[string]string)
	book["title"] = "buku novel"
	book["author"] = "Rakha"
	book["hilang"] = "salah"

	fmt.Println(book)
	delete(book, "hilang")
	fmt.Println(book)
}
