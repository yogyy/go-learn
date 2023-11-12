package main

import "fmt"

func main() {
	names := []string{"Batman", "Robin", "Nightwing"}
	for i := 0; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// or

	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}

	// or

	for _, v := range names {
		fmt.Println(v)
	}
}
