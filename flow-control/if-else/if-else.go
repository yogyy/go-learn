package main

import "fmt"

func main() {
	name := "Barry"

	if name == "John" {
		fmt.Println("hello John")
	} else if name == "Barry" {
		fmt.Println("hello Barry")
	} else {
		fmt.Println("you not John")
	}

	if length := len(name); length > 8 {
		fmt.Println("nama terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}

}
