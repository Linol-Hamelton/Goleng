package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Scan(&str)

	maxDigit := 0
	for i := 0; i < len(str); i++ {
		digit, _ := strconv.Atoi(string(str[i]))
		if digit > maxDigit {
			maxDigit = digit
		}
	}

	fmt.Println(maxDigit)
}
