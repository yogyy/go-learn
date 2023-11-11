package main

import "fmt"

func getFullName() (string, string) {
	return "John", "Constantine"
}

func getCompleteName() (firstName, middleName, lastName string) {
	firstName = "Bruce"
	middleName = "Gotham"
	lastName = "Wayne"

	return firstName, middleName, lastName
}

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func getGoodBye(name string) string {
	return "Good Bye " + name
}

type Filter func(string) string

func sayHiwFilter(name string, filter Filter) {
	filtered := filter(name)

	fmt.Println(filtered)
}

func filterName(name string) string {
	if name == "Barry" {
		return "Flash"
	} else {
		return name
	}
}

type Blacklist func(string) bool

func registerUser(name string, blacklist Blacklist) {
	if blacklist(name) {
		fmt.Println("You are not Allowed", name)
	} else {
		fmt.Println("Welcome!", name)
	}
}

func factorialLoop(value int) int {
	res := 1
	for i := value; i > 0; i-- {
		res *= i
	}

	return res
}

func factorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * factorialRecursive((value - 1))
	}
}

func main() {
	counter := 0

	increment := func() {
		fmt.Println("Increment")
		counter++
	}

	increment()
	increment()
	fmt.Println(counter)
}
