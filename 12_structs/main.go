package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	// firstName string
	// lastName  string
	// city      string
	// gender    string
	firstName, lastName, city, gender string
	age       int
}

// greeting method (value receiver)
func (p Person) greet() string {
	return "Hi " + p.firstName + " " + p.lastName + "! According to our records you're a " + strconv.Itoa(p.age) + " year old " + p.gender + " living in " + p.city + "."
}

//hasBirthday method (pointer receiver )
func (p *Person) hasBirthday() {
	p.age++
}

//getMarried method (pointer receiver)

func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "male" {
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	person1 := Person{firstName: "Joey", lastName: "Douglas", city: "Austin", gender: "male", age: 42}
	person2 := Person{"Cindy", "Garcia", "Austin", "female", 36}
	person1.hasBirthday()
	person1.getMarried("Garcia")
	person2.getMarried("Douglas")
	fmt.Println(person1.greet())
	fmt.Println(person2.greet())

}
