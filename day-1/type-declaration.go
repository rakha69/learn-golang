package main

import "fmt"

func main() {
	type NoKTP string

	var ktpRakha NoKTP = "241011400043"

	var contoh string = "231011400043"
	var contohKtp NoKTP = NoKTP(contoh)

	fmt.Println(ktpRakha)
	fmt.Println(contohKtp)

}
