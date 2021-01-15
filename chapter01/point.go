package main

import "fmt"

func main() {
	A := 2

	//指针
	//C := &A

	var c *int = &A

	//fmt.Printf("%T\n %T\n", C, c)
	//fmt.Printf("%d %d\n", C, c)

	//修改原始值
	*c = 4
	fmt.Println(A)
}
