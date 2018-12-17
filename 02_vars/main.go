package main

import "fmt"

func main() {

	name, email := "Joey", "joey.douglas@gmail.com"
	age := 42

	fmt.Println(name, age, email)
	fmt.Printf("%T\n", age)
}
