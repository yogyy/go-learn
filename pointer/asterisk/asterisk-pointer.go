package main

import "fmt"

type Address struct {
	Place, City, Earth string
}

func main() {
	address1 := Address{"The Batcave", "outside Gotham City", "Earth Prime"}
	address2 := &address1
	address2.Place = "Wayne Manor"

	*address2 = Address{"S.T.A.R. Labs", "Central City", "Earth Prime"}
	fmt.Println(address1)
	fmt.Println(address2)
}
