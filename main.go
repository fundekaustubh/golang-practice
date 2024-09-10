package main

import (
	"fmt"
	"strconv"
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

var people = []Person{}

func getPeople() {
	for _, person := range people {
		fmt.Println(person.getFirstName(), person.getLastName())
	}
}


func main() {

	for i := 1; i <= 10; i++ {
		people = append(people, Person{
			firstName: "FirstName " + strconv.Itoa(i),
			lastName: "LastName " + strconv.Itoa(i),
		})
	}

	fmt.Println("All people: ")
	getPeople()
}