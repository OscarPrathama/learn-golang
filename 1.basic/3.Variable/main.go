/*
-> ada 2 jenis variable dalam go, yakni yang tipe datanya di definisikan dan yang tidak
*/

package main

import "fmt"

func _break() {
	fmt.Println("=====================================")
	fmt.Println()
}

func main() {

	// variable definition
	/*---------------------------------------*/
	// var yang di definisikan tipe datanya
	var firstName string = "Christiano"
	var lastName string
	lastName = "Ronaldo"

	fmt.Printf("hello %s %s \n", firstName, lastName)
	_break()

	// var yg tidak di definisikan tipe datanya
	var country = "Indonesia"
	age := 25
	// tanda := hanya bsa di definisikan 1x, u/ selanjutnya harus menggunaka = (age = 25)
	// tanda := hanya bsa di definisikan di dalam block

	fmt.Printf("Hello saya %s %s, saya tinggal di %s, usia saya %d", firstName, lastName, country, age)

	// deklarasi multivariable
	/*---------------------------------------*/
	var fruit, vehicle string
	var number int
	fruit, number, vehicle = "apple", 241, "car"
	fmt.Println(fruit, number, vehicle)
	_break()
	// atau

	var state, city, province string = "Indonesia", "Depok", "Jawa Barat"
	fmt.Println(state, city, province)
	_break()
	// atau

	color, flower, animal := "blue", "rose", "butterfly"
	fmt.Println(color, flower, animal)
	_break()
	// atau

	_number, _string, _bool, _decimal := 1, "hello", true, 9.5
	fmt.Println(_number, _string, _bool, _decimal)
	_break()

	// var underscore
	/*---------------------------------------*/
	/*
		dalam go, semua var harus digunakan, lalu bgmn caranya membiarkan var yg mengannggur ??
		-> jwbnnya adalah, kita bisa menggunakan var underscore (_)
		-> isi dlm var underscore tdk dapat ditampilkan
	*/
	_ = "tidak berguna"

	// mendefinisikan var dgn keyword new
	/*---------------------------------------*/
	/*
		-> keyword new digunakan u/ membuat var pointer
		-> data pointer merupakan data alamat memori
	*/
	address := new(string)
	fmt.Println(address)
}
