package main

import (
	"fmt"
)

func main() {
	names := [...]string{"rakha", "kiwil", "ucup", "budi", "qila", "alim"}

	slice1 := names[4:6]
	fmt.Println(slice1)

	slice2 := names[0:5]
	fmt.Println(slice2)

	slice3 := names[3:]
	fmt.Println(slice3)

	slice4 := names[:]
	fmt.Println(slice4)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}

	daySlice1 := days[5:] // sabtu, minggu
	fmt.Println(daySlice1)

	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu Baru"
	fmt.Println(daySlice1)
	fmt.Println(days)

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "sabtu lama"
	//days := [...]string{"senin", "selasa", "rabu", "ksmis", "jumat", "sabtu", "minggu" "libur baru"}
	fmt.Println(daySlice1)
	fmt.Println(daySlice2)
	fmt.Println(days)

	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Rakha"
	newSlice[1] = "Rakha"
	// newslice2 = "Rakha" // error, harusmya menggunakan append

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newslice2 := append(newSlice, "Rakha")
	fmt.Println(newslice2)
	fmt.Println(len(newslice2))
	fmt.Println(cap(newslice2))

	newslice2[0] = "capi"
	fmt.Println(newslice2)
	fmt.Println(newSlice)

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)

	//perbedaan array dan slice
	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
