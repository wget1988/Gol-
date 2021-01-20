package main

import "fmt"

func returnTest() int {
	return 0
	return 1
}

func calc2(a, b int) (sum int, diff int, product int, merchant int) {
	sum = a + b
	diff = a - b
	product = a * b
	merchant = a / b

	return
}
func main() {
	fmt.Println(returnTest())

	fmt.Println(calc2(5, 2))
}
