package main

import (
	"fmt"
	"strconv"
)

// Define Person struct
type Person struct {
	firstName, lastName, city, gender string
	age                               int
}

// Greeting method (value receiver)
func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age)
}

// hasBirthday method (pointer receiver)
func (p *Person) hasBirthday() {
	p.age++
}

// getMarried (pointer receiver)
func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "m" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	// Init person using struct
	person1 := Person{
		firstName: "Cristian",
		lastName:  "Silva",
		city:      "Joinville",
		gender:    "Male",
		age:       19,
	}

	person2 := Person{
		"Ana",
		"Car",
		"Joinville",
		"f",
		19,
	}

	fmt.Println(person1, person2)
	fmt.Println(person1.firstName)

	person1.age = 20

	fmt.Println(person1)
	person1.hasBirthday()
	person2.getMarried(person1.lastName)

	fmt.Println(person1.greet())
	fmt.Println(person2)
}
