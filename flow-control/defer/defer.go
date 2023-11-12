package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApp() {
	defer logging()
	fmt.Println("Run application")
}

func main() {
	runApp()
}
