package main

import (
	"fmt"
	"regexp"
)

func main() {

	var password string

	fmt.Scanln(&password)

	pattern := "^[A-Za-z0-9]+$"

	regex := regexp.MustCompile(pattern)

	if regex.MatchString(password) == true && len(password) >= 5 {
		fmt.Println("Ok")
	} else {
		fmt.Println("Wrong password")
	}
}
