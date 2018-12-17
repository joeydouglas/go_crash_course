package main

import "fmt"

func main() {
	x := 50
	y := 10

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else {
		fmt.Printf("%d is less than %d\n", y, x)
	}

	color := "pink"

	if color == "red" {
		fmt.Println("Color is RED")
	} else if color == "blue" {
		fmt.Println("Color is BLUE")
	} else {
		fmt.Println("Color is neither RED nor BLUE")
	}

}
