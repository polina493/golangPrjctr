package main

import "fmt"

func choice(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

func main() {
	var answer int
	fmt.Println("Enter int 1 or 2:")
	fmt.Scan(&answer)

	result := choice(answer)

	if result {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
