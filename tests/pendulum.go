package main

import (
	"fmt"
	"math"
)

func T() float64 {
	w := W()
	return 6 / w
}

func W(k int) int {
	m := M()
	return math.Sqrt(k / m)
}

func M(p, v int) int {
	return p * v
}

func main() {
	t := T()
	fmt.Println(math.Round(t))
}
