package main

import "fmt"

func main() {

	fmt.Println("main start")

	panic("err")

	fmt.Println("over")
}
