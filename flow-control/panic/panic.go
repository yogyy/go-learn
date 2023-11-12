package main

import "fmt"

func endApp() {
	fmt.Println("App End")
}

func runApp(error bool) {
	defer endApp()

	if error {
		panic("ERROR")
	}

}

func main() {
	runApp(true)
}
