package main

import "fmt"

func main() {
	var country [3]string
	country[0] = "England"
	country[1] = "France"
	country[2] = "Germany"

	var cities = [3]string{
		"London",
		"Paris",
		"Berlin",
	}

	fmt.Println(country[0])
	fmt.Println(country)
	fmt.Println(cities)
}

/*
1. Saat membuat Arraynya, kita harus menentukan jumlah data yang bisa ditampung array tersebut.
2. Jumlah data yang bisa ditampung oleh array, tidak bisa ditambah lagi.
*/
