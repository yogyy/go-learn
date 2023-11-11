package main

import "fmt"

func main() {
	var newSlice []string = make([]string, 2, 5)
	newSlice[0] = "Bruce"
	newSlice[1] = "Wayne"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "and Alfred")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	newSlice2[0] = "Thomas"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	iniArray := [...]int{1, 2, 4, 5}
	iniSlice := []int{1, 2, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
