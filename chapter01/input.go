package main

import "fmt"

func main() {

	var name string
	fmt.Print("请输入姓名：")
	fmt.Scan(&name)
	fmt.Println("姓名：", name)

	var age int
	fmt.Print("请输入年龄：")
	fmt.Scan(&age)
	fmt.Println("年龄：", age)
}
