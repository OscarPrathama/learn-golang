package main

import "fmt"

func main() {
	// "Simple loop"
	for i := 0; i < 4; i++ {
		fmt.Println("Perulangan ke : ", i)
	}
	fmt.Println("===========================")

	// nested loop
	for y := 0; y < 5; y++ {
		for z := 0; z < y; z++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	fmt.Println("===========================")

	// array loop
	cities := []string{
		"jakarta", "medan", "surabaya", "banten",
	}
	for _, v := range cities {
		fmt.Println(v)
	}
	fmt.Println("===========================")

	// map loop
	teritory := make(map[string]string)
	teritory["country"] = "Indonesia"
	teritory["city"] = "Jakarta"
	teritory["district"] = "Tanah Abang"
	for k, v := range teritory {
		fmt.Println(k, "=", v)
	}
}
