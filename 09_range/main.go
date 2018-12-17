package main

import "fmt"

func main() {

	ids := []int{5, 10, 15, 20, 25, 30}

	// for i, id := range ids {
	// 	fmt.Printf("%d - ID:%d\n", i, id)
	// }

	for _, id := range ids {
		fmt.Printf("ID:%d\n", id)
	}

	sum := 0
	for _, id := range ids {
		sum += id
	}
	fmt.Println("Sum =", sum)

	emails := map[string]string{"Bob": "bob@gmail.com", "Jill": "jill@gmail.com", "Bart": "bart@gmail.com"}

	for key, value := range emails {
		fmt.Printf("%s:%s\n", key, value)
	}

}
