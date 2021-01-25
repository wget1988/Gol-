package main

import "fmt"

func changeInt(a int) {
	a = 100
}

func changeSlice(s []int) {

	s[0] = 100
}

func changePoint(p *int) {
	*p = 88
}
func main() {
	num := 1
	changeInt(num) //值不变

	nums := []int{1, 2, 3}
	changeSlice(nums) //值修改

	num1 := 100
	changePoint(&num1)
	fmt.Println(num1)
}
