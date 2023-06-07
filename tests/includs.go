package main

import (
	"fmt"
	"strings"
)

func main() {

	var X, S string

	fmt.Scanln(&X)
	fmt.Scanln(&S)

	index := strings.Index(X, S)

	if index != -1 {
		fmt.Println(index)
	} else {
		fmt.Println(-1)
	}
}
