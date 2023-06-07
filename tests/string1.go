package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func main() {

	text, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	firstChar := []rune(text)[0]

	if unicode.IsUpper(firstChar) && text[len(text)-1] == '.' {
		fmt.Println("Right")
	} else {
		fmt.Println("Wrong")
	}
}
