package main

import "fmt"

func main() {

	var word, result string

	fmt.Scan(&word)

	for i := 0; i < len(word); i++ {
		if i%2 == 1 {
			result += string(word[i])
		}
	}
	fmt.Println(result)
}
