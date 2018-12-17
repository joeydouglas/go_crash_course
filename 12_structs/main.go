package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	city      string
	gender    string
	age       int
}

func (p Person) greet() string {
	return "Hi " + p.firstName + " " + p.lastName + "!! According to our records you're a " + strconv.Itoa(p.age) + " year old " + p.gender + " living in " p.city + "."
}

func main() {
	// person1 := Person{firstName: "Samantha", lastName: "Mathis", city: "Austin", gender: "shemale", age: 300}
	person1 := Person{"Joey", "Douglas", "Austin", "male", 42}
	fmt.Println(person1.greet())
}
