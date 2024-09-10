package main

import (
	"fmt"
)

type Person struct {
	firstName	string
	lastName	string
}

func (p Person) getFirstName() string {
	return p.firstName
}

func (p Person) getLastName() string {
	return p.lastName
}

func main() {
	p := Person{
		firstName: "Kaustubh",
		lastName: "Funde",
	}
	fmt.Println("This person is", p.getFirstName(), p.getLastName())
}