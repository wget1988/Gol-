package main

import "fmt"

func main() {

	var age = 30

	age = age + 1
	fmt.Println("明年：", age)

	age = age + 1

	fmt.Println("后年：", age) //默认输出一个换行

	fmt.Print("打印内容") //不加ln不打印换行
	fmt.Printf("%s", "age") 
}
