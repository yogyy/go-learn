package main

import "fmt"

func main() {
	type noKTP = string
	var ktpJohn = "200999"

	var contoh string = "200666"
	var contohKtp noKTP = noKTP(contoh)

	fmt.Println(ktpJohn)
	fmt.Println(contohKtp)
}
