package main

import (
	"fmt"
)

func main() {

	// if statement
	month := "Maret"
	if month == "April" {
		fmt.Printf("Sekarang bulan : %s \n", month)
	} else {
		fmt.Printf("Sekarang bukan bulan April \n")
	}
	fmt.Println("===================")

	/*
		var temporary in if else statement
		-> hny bsa di gunakan dalam blok itu saja
		-> pendefinisian var temporary harus menggunakan tanda :=
		-> pendefinisian var temporary tidak boleh menggunakan keyword var
	*/
	var mtk, bhs_indo, bhs_ing = 60, 70, 74
	if avg := (mtk + bhs_indo + bhs_ing) / 3; avg > 90 && avg <= 100 {
		fmt.Printf("Nilai anda : %d, Great !! \n", avg)
	} else if avg >= 60 && avg <= 90 {
		fmt.Printf("Nilai anda : %d, Average \n", avg)
	} else {
		fmt.Printf("Nilai anda : %d, Bad \n", avg)
	}
	fmt.Println("===================")

	// switch statement
	var color = "blue"
	switch color {
	case "blue":
		fmt.Printf("it's blue \n")
	case "red":
		fmt.Printf("it's red \n")
	default:
		fmt.Printf("i don't know \n")
	}
	fmt.Println("===================")

	// switch case multiple condition
	var car = "lamborgini"
	switch car {
	case "nissan":
		fmt.Printf("i'ts nissan \n")
	case "ferari", "lamborgini", "toyota":
		fmt.Printf("maybe one of the list above \n")
	default:
		fmt.Printf("i don't know what is car it is... \n")
	}
	fmt.Println("===================")

	// kurung kurawal pada keyword case & default
	/*
		-> metode ini bagus u/ kondisi yg di dalamnya trdpt bnyk statement
	*/
	a := 30
	switch a {
	case 100:
		fmt.Printf("Good \n")
	case 70, 80, 90:
		fmt.Printf("Average \n")
	default:
		{
			fmt.Printf("No problem \n")
			fmt.Printf("Keep trying \n")
		}
	}
	fmt.Println("===================")

	// switch dgn gaya if...else
	b := 86
	switch {
	case b > 90 && b <= 100:
		fmt.Printf("Good bro \n")
	case b >= 70 && b < 90:
		fmt.Printf("Average bro \n")
	default:
		{
			fmt.Printf("No problem bro \n")
			fmt.Printf("Keep trying bro \n")
		}
	}
	fmt.Println("===================")

	// switch dgn keyword fallthrough
	/*
		-> pada pemrograman lain ketika case sudah terpenuhi maka pengecekan kondisi tidak akan diteruskan
		-> namun pada go, terdapat keyword fallthrough yang digunakan untuk melanjutkan proses pengecekan pada case selanjutnya
	*/
	var c = 50
	switch {
	case c >= 70 && c <= 100:
		fmt.Printf("good \n")
	case c >= 40 && c <= 69:
		fmt.Printf("average \n")
		fallthrough
	case c <= 60:
		fmt.Printf("in this value, you need learn more \n")
	default:
		fmt.Printf("no problem just keep going \n")
	}
	fmt.Println("===================")
}
