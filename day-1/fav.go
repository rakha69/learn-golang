package main

import "fmt"

func getGoodBye(name string) string {
	return "good bye" + name
}

func main() {
	contoh := getGoodBye
	misal := getGoodBye

	fmt.Println(contoh("Rakha"))
	fmt.Println(misal("Rakha"))

}
