package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n int
	fmt.Scan(&n)

	s := strconv.Itoa(n)
	var result string
	for _, c := range s {
		digit, _ := strconv.Atoi(string(c))
		result += strconv.Itoa(digit * digit)
	}

	fmt.Println(result)
}
