package main

import "fmt"

func main() {
	// tipe data numerik positif
	var angka1 uint8 = 255
	var angka2 uint16 = 65535
	var angka3 int8 = 127
	fmt.Println(angka1, angka2, angka3)
	fmt.Println("========================")

	// tipe data numerik non desimal
	var angka4 = 12.53
	fmt.Println()
	fmt.Printf("%f\n", angka4)
	fmt.Printf("%.2f\n", angka4)
	fmt.Println("========================")

	// tipe data bool
	var isset bool = true
	fmt.Println()
	fmt.Printf("isset ? %t\n", isset)
	fmt.Println("========================")

	// tipe data string
	fmt.Println()
	var output string = "hello world."
	output2 := `witing tresno.
	jalaran soko kulino`
	fmt.Printf("output : %s %s\n", output, output2)
	fmt.Println("========================")

	// nil. nil bukan tipe data, tapi sebuah nilai kosong
	/*
		semua tipe data memilki zero value
		1. zero value strong adalah ""
		2. zero value bool adalah false
		3. zero value tipe numerik non desimal adalah 0
		4. zero value tipe numerik desimal adalah 0.0
	*/
}
