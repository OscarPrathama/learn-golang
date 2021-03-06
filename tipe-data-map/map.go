package main

import "fmt"

func main() {
	teritory := map[string]string{
		"country":  "indonesia",
		"city":     "jakarta",
		"district": "tanah abang",
	}

	fmt.Println(teritory)
	fmt.Println("======================")
	fmt.Println(teritory["city"])
	fmt.Println("======================")
	for k, v := range teritory {
		fmt.Println(k, "=", v)
	}
}
