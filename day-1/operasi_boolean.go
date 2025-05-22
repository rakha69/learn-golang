package main

import "fmt"

func main() {
	var nilaiAkhir = 98
	var absensi = 80

	var LuLusNilaiAkhir bool = nilaiAkhir > 75
	var LuLusAbsensi bool = absensi > 78

	var lulus bool = LuLusNilaiAkhir && LuLusAbsensi

	fmt.Println(lulus)
}
