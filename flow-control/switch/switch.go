package main

import "fmt"

func main() {
	name := "Flash"

	switch name {
	case "Batman":
		fmt.Println("hello Batman")
	case "Superman":
		fmt.Println("hello Superman")
	case "Flash":
		fmt.Println("hello Flash")
	default:
		fmt.Println("who are you?")
	}

	switch length := len(name); length > 8 {
	case true:
		fmt.Println("nama terlalu panjang")
	case false:
		fmt.Println("nama sudah benar")
	}

	length := len("superman")
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
