package main

import "fmt"

func main() {
	// var names [3]string
	// names[0] = "John"
	// names[1] = "Constantine"
	// names[2] = "Prime"

	// fmt.Println(names[0], names[1])

	// var earths = [3]string{
	// 	"Earth-Zero",
	// 	"Earth-Prime",
	// 	"Earth-One",
	// }

	// fmt.Println(earths)
	// fmt.Println(len(earths))

	batmanVillains := [...]string{"Joker", "Ra's al Ghul", "Bane", "Riddler", "Penguin", "Two-Face", "Scarecrow"}

	slice1 := batmanVillains[4:6]
	fmt.Println(slice1) // [Penguin Two-Face]

	slice2 := batmanVillains[:2]
	fmt.Println(slice2) // [Joker Ra's al Ghul]

	slice3 := batmanVillains[3:]
	fmt.Println(slice3) // [Riddler Penguin Two-Face Scarecrow]

	slice4 := batmanVillains[:]
	fmt.Println(slice4) // [Joker Ra's al Ghul Bane Riddler Penguin Two-Face Scarecrow]

	fromSlice := batmanVillains[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)
}
