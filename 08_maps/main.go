package main

import "fmt"

func main() {
	//Define map
	// emails := make(map[string]string)
	//
	// emails["Bob"] = "bob@gmail.com"
	// emails["Jill"] = "jill@gmail.com"
	// emails["Bart"] = "bart@gmail.com"

	emails := map[string]string{"Bob": "bob@gmail.com", "Jill": "jill@gmail.com", "Bart": "bart@gmail.com"}

	fmt.Println(emails["Bob"])

	//Delete map
	delete(emails, "Bob")
	fmt.Println(emails)

}
