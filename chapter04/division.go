package main

import (
	"errors"
	"fmt"
)

func division(a, b int) (int, error) {

	if b == 0 {
		return -1, errors.New("division is by zero")
	}
	return a / b, nil
}

func main() {

	fmt.Println(division(1, 3))

	if v, err := division(1, 0); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}
	e := fmt.Errorf("Error:%s", "errors")
	fmt.Printf("%T,%v\n", e, e)
}
