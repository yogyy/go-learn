package main

import "fmt"

type Address struct {
	Place, City, Earth string
}

func main() {
	// address1 := Address{"The Batcave", "outside Gotham City", "Earth Prime"}
	// address2 := address1

	// address2.Place = "S.T.A.R. Labs"
	// address2.City = "Central City"
	// fmt.Println(address1)
	// fmt.Println(address2)

	address1 := Address{"The Batcave", "outside Gotham City", "Earth Prime"}
	address2 := &address1
	fmt.Println(address1)

	address2.Earth = "Earth One"
	fmt.Println(address1)
	fmt.Println(address2)
}
