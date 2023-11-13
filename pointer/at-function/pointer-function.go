package main

import "fmt"

type Address struct {
	Place, City, Earth string
}

func ChangeAddressToGotham(address *Address) {
	address.City = "Gotham"
}

func main() {
	address := Address{}
	ChangeAddressToGotham(&address)
	fmt.Println(address)
}
