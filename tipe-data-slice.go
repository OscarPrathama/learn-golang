package main

import "fmt"

func main() {
	var nilai_array = [...]string{
		"satu", "dua", "tiga", "empat", "lima", "enam", "tujuh", "delapan",
	}

	var nilai_slice_1 = nilai_array[3:5]

	fmt.Println("nilai_slice_1")
}

/*
1. Tipe data slice mengambil nilai pada array
2. Jika nilai pada slice dirubah, maka nilai pada array juga ikut terubah
3. Namun berbeda dengan array, pada tipe data slice anda dapat menambahkan jumlah datanya
*/
