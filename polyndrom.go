package main

import (
	"fmt"
)

func main() {

	var input string

	fmt.Scanln(&input)

	runes := []rune(input)
	length := len(runes)

	for i := 0; i < length/2; i++ {
		runes[i], runes[length-1-i] = runes[length-1-i], runes[i]
	}

	if string(runes) == string(input) {
		fmt.Println("Палиндром")
	} else {
		fmt.Println("Нет")
	}
}
