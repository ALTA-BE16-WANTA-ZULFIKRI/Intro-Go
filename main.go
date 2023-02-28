package main

import "fmt"

func main() {

	fmt.Println("Hello World")
	fmt.Println("wanta zulfikri")
	fmt.Println("zulfikri")
	fmt.Println("wanta")

	var umur = 20
	var nama = "jerry"
	var bolehtidak = false // boolean
	// var uang = 1000
	var pi = 3.14

	fmt.Println(umur, nama, bolehtidak, pi)

	nama1 := "wanta"
	fmt.Println(nama1)
	// var uang = 1000
	var uang1 int = 1000
	var angkapositif uint = 100
	var decimal float32 = 4.11
	var ijin bool = true
	fmt.Println(uang1, angkapositif, decimal, ijin)

	alamat, umur3 := "pasuruan", 50
	fmt.Println(alamat, umur3)

	totalbayar := uang1 * int(angkapositif)
	fmt.Println(totalbayar)

	jumlahumur := umur3 + int(decimal)
	fmt.Println(jumlahumur)

	jumlahumuranda := float32(umur3) + decimal
	fmt.Println(jumlahumuranda)

	var umur4 int
	umur5 := 0
	fmt.Println(umur4, umur5)

	gabunganstr := nama + alamat
	fmt.Println(gabunganstr)

	gabunganstr2 := fmt.Sprint(nama, " ", alamat)
	fmt.Println(gabunganstr2)

	fmt.Println("Kondisi awal", umur)
	umur = umur + 1
	fmt.Println("setelah di +1", umur)
	umur++
	fmt.Println("setelah di ++", umur)

	umur += 2
	fmt.Println("setelah di +=2", umur)

	umur += 2
	fmt.Println("setelah di +=2", umur)

	umur /= 2
	fmt.Println("setelah di /=2", umur)

	fmt.Scanln(&umur)

	fmt.Println("umur hasil input", umur)

	fmt.Scanln(&alamat)
	fmt.Println("alamat hasil input", alamat)

	alamat = "12345678"
	umur = angka
	var input int
	fmt.Println(umur)

}
