package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func (customer Customer) sayHello(name string) {
	fmt.Println("Hello", name, "my name is", customer.Name)
}

func main() {
	var john Customer
	john.Address = "Earth Prime"
	john.Name = "Constantine"
	john.Age = 20

	fmt.Println(john.Name)

	bruce := Customer{
		Name:    "Batman",
		Address: "Gotham",
		Age:     48,
	}

	fmt.Println(bruce)

	bruce.sayHello("john")
}
