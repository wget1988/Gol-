package main

import "fmt"

func main() {

	//	name := "sql"

	add2 := func(n int) int {
		return n + 2
	}

	fmt.Println(add2(2))

	addBase := func(base int) func(int) int {
		//返回一个函数
		return func(n int) int {
			return base + n
		}
	}

	add100 := addBase(100)
	fmt.Println(add100(100))
}
