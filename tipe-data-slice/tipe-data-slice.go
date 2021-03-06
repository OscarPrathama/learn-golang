package main

import "fmt"

func main() {

	var slice = make([]string, 2, 5)
	slice[0] = "Hello"
	slice[1] = "world"

	fmt.Println(slice)

	//--------------------------------------//
	var nilaiArray = [...]string{
		"satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan",
	}

	var nilaiSlice1 = nilaiArray[2:6]

	fmt.Println(nilaiSlice1)
}

/*
1. Tipe data slice mengambil nilai pada array
2. Jika nilai pada slice dirubah, maka nilai pada array juga ikut terubah
3. Namun berbeda dengan array, pada tipe data slice anda dapat menambahkan jumlah datanya
*/
