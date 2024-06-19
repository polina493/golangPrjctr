package main

import (
	"fmt"
	"strings"
)

func search(s string, t []string) []string {
	var result []string
	for _, word := range t {
		if strings.Contains(word, s) {
			result = append(result, word)
		}
	}
	return result
}

func main() {
	text := []string{"some text here, some other text here", "also some text here", "just another line"}
	fmt.Println("Enter a word that you want to find:")
	var word string
	fmt.Scan(&word)

	result := search(word, text)

	fmt.Println("Result:")
	if len(result) == 0 {
		fmt.Println("There is no such word.")
	} else {
		for _, word := range result {
			fmt.Println(word)
		}
	}
}
