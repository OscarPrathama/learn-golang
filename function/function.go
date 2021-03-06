package main

import "fmt"

func breakRule() {
	fmt.Println("=================================")
	fmt.Println()
}

// simple function
func hello() {
	fmt.Println("hello world")
}

// function with parameter
func tambah(a int, b int) {
	fmt.Println(a, "+", b, "=", a+b)
}

// function with return value
func biodata(name string) string {
	if name == "" {
		return "insert your name"
	}
	return "Your name is " + name
}

// returning multiple value
func getInfo(name string, age string, address string) (string, string, string) {
	return "your name :" + name, ", your age :" + age, ", your address :" + address
}

// returning multiple value 2
func getInfo2(name string, age string, address string) (string, string, string) {
	return name, age, address
}

// named return value
func getInfo3() (name string, fee int) {
	name = "Hasan"
	fee = 900000
	return
}

// variadic function
func totalPayment(prices ...int) int {
	subtotal := 0
	for _, price := range prices {
		subtotal += price
	}
	return subtotal
}

func main() {
	// simple function
	hello()
	breakRule()

	// function with parameters
	tambah(12, 32)
	tambah(44, 53)
	breakRule()

	// function with return value
	fmt.Println(biodata("Yugi"))
	breakRule()

	// returning multiple value
	fmt.Println(getInfo("rudi", "21", "berlin"))
	breakRule()

	// returning multiple value 2
	myname, myage, myaddress := getInfo2("habibi", "44", "paris")
	fmt.Println(myname, myage, myaddress)
	breakRule()

	// named return value
	employeeName, employeeFee := getInfo3()
	fmt.Println(employeeName)
	fmt.Println(employeeFee)
	breakRule()

	// variadic function
	subTotal := totalPayment(25000, 1500, 10000) //price, tax, shipping
	fmt.Println("Jumlah yang harus anda bayar Rp", subTotal)
}
