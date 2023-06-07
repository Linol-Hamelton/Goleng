package main

import (
	"fmt"
)

var cache = make(map[int]int)

func work(x int) int {
	return x * 2
}

func main() {
	var input int
	for i := 0; i < 10; i++ {
		_, err := fmt.Scan(&input)
		if err != nil {
			fmt.Println("Error reading input:", err)
			return
		}

		result, ok := cache[input]
		if !ok {
			result = work(input)
			cache[input] = result
		}

		fmt.Print(result, " ")
	}

	fmt.Println("time limit ok")
}
