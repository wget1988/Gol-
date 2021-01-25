package main

import "fmt"

func main() {
	defer func() { //延迟执行，当函数退出时候执行
		fmt.Println("deffer1")
	}()

	defer func() { //延迟执行，当函数退出时候执行
		fmt.Println("deffer2") //defer2先执行，defer1后执行先进后出，堆栈应用
	}()

	fmt.Println("main.over")
}
