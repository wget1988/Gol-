package main

import "fmt"

func tower(a, b, c string, layer int) {
	if layer == 1 {
		fmt.Println(a, "->", c)
		return
	}
	//n-1 个a的借助c挪到b
	tower(a, c, b, layer-1)
	fmt.Println(a, "->", c)
	//b 借助a移动到c
	tower(b, a, c, layer-1)
}

func main() {

	fmt.Println("三层")
	tower("A", "B", "C", 3)
}
