package main

import (
	"fmt"
)

func removeDuplicates(input string) string {
	var result []rune

	for _, char := range input {
		if !isDuplicate(result, char) {
			result = append(result, char)
		}
	}

	return string(result)
}

func isDuplicate(slice []rune, char rune) bool {
	for _, c := range slice {
		if c == char {
			return true
		}
	}
	return false
}

func main() {
	var input string
	fmt.Scanln(&input)
	result := removeDuplicates(input)
	fmt.Println(result)
}
