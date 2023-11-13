package main

import "fmt"

func random() any {
	return "OK"
}

func main() {
	var res any = random()
	// var resString string = res.(string)

	// fmt.Println(resString)

	switch value := res.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown", value)
	}
}
