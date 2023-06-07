package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   []string{"A", "B"},
		100:  []string{"C", "D"},
		1000: []string{"E", "F"},
	}

	cityPopulation := map[string]int{
		"A": 10,
		"C": 200,
		"D": 300,
		"E": 500,
		"G": 700,
	}

	for _, city := range groupCity[100] {
		_, exists := cityPopulation[city]
		if !exists {
			delete(cityPopulation, city)
		}
	}

	// Print the fixed cityPopulation map
	for city, population := range cityPopulation {
		fmt.Println(city, population)
	}
}
