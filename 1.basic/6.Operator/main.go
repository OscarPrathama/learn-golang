package main

import "fmt"

func main() {
	// /aritmatika
	tambah := 12 + 12
	kurang := 12 - 12
	kali := 12 * 12
	bagi := 12 / 12
	modulus := 12 % 12
	fmt.Println("tambah : ", tambah)
	fmt.Println("kurang : ", kurang)
	fmt.Println("kali : ", kali)
	fmt.Println("bagi : ", bagi)
	fmt.Println("modulus : ", modulus)
	fmt.Println("==========================")

	// logika
	x := true
	y := false

	xAndY := x && y
	xOrY := x || y
	negasiY := !y

	fmt.Printf("x && y \t(%t) \n", xAndY)
	fmt.Printf("x || y \t(%t) \n", xOrY)
	fmt.Printf("!y \t(%t) \n", negasiY)
	// \t digunakan sebagai tabulasi untuk merapikan output pada console
	fmt.Println("==========================")

	// perbandingan
	a := 12
	b := (a == 13)
	fmt.Printf("%t", b)
}
