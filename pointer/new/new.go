package main

import "fmt"

type Address struct {
	Place, City, Earth string
}

func main() {
	var alamat1 = new(Address)
	var alamat2 = alamat1

	alamat2.City = "Star City"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
}
