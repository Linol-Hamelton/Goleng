package main

import (
	"fmt"
)

func addAsterisks(s string) string {

	result := []rune{}

	for i, ch := range s {
		result = append(result, ch)

		if i < len(s)-1 && s[i+1] != ' ' {
			result = append(result, '*')
		}
	}

	return string(result)
}

func main() {
	var input string
	fmt.Scanln(&input)

	output := addAsterisks(input)

	fmt.Println(output)
}
