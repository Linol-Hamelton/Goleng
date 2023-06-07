package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	err := errors.New("ошибка")
	if b == 0 {
		return 0, err
	} else {
		return a / b, nil
	}
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	result, err := divide(a, b)
	if err == nil {
		fmt.Println(result)
	} else {
		fmt.Println(err)
	}
}
