package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Lucy",
		"address": "Earth-Prime",
	}

	fmt.Println(person["name"])
	fmt.Println(person["address"])

	character := make(map[string]string)
	character["name"] = "Flash"
	character["power"] = "Super Speed"
	character["main-villain"] = "Reverse Flash"

	fmt.Println(character)
	delete(character, "main-villain")
	fmt.Println(character)
}
