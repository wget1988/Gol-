package main

import "fmt"

//函数定义,无参，无返回值
func sayHelloWord() {

	fmt.Println("hello world")
}

func sayHi(name string) {

	fmt.Println("hello", name)
}

//带返回值
func add(a int, b int) int {

	return a + b
}

//可变参数 args 但类型必须一致

func addN(a, b int, args ...int) int {
	fmt.Println(a, b, args)
	fmt.Printf("%T\n", args) //切片类型

	total := a + b

	for _, v := range args {
		total += v
	}
	return total
}

//计算器
func calc(op string, a, b int, args ...int) int {

	switch op {
	case "add":
		return addN(a, b, args...)
	}
	return 0
}

func print(callback func(...string), args ...string) {
	fmt.Println("输出")
	callback(args...)
}

func list(args ...string) {
	for i, v := range args {
		fmt.Println(i, v)
	}
}

func main() {

	sayHelloWord()
	sayHi("sql")

	name := "hy"
	sayHi(name)

	num1 := 1
	num2 := 2
	rtNum := add(num1, num2)
	fmt.Println(rtNum)

	fmt.Println(addN(1, 2, 3))

	nums := []int{1, 3, 5, 8}

	fmt.Println(nums[:1])
	fmt.Println(nums[2:])

	//nums = append(nums[:1], 5, 8)
	nums = append(nums[:1], nums[2:]...)
	fmt.Println(nums)

	//函数赋值给变量
	var f func(int, int) int = add
	fmt.Println("%T\n", f)
	fmt.Println(f(1, 4))

	print(list, "A", "C")

	//匿名函数 在函数内定义一个变量等于一个变量，直接调用就不用取原函数名字
	sayHello := func(name string) {
		fmt.Println("hello", name)
	}

	sayHello("h")

	func(name string) {
		fmt.Println("hi", name)

	}("kk")

	//两种方式
	value := func(args ...string) {
		for _, v := range args {
			fmt.Println(v)
		}
	}

	print(value, "A", "B", "C")

	print(func(args ...string) {
		for _, v := range args {
			fmt.Println(v, "\t")
		}
		fmt.Println()
	}, "A", "B", "C")

}
