package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b float64
	fmt.Scan(&a, &b)
	c := math.Sqrt((math.Pow(a, 2)) + math.Pow(b, 2))

	fmt.Println(c)
}
