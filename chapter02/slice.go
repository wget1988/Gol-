package main

import "fmt"

func main() {

	var nums []int

	// fmt.Printf("%T\n", nums)
	// fmt.Println(nums == nil)

	// //字面量
	// nums = []int{1, 3, 5, 7}
	// fmt.Println(nums)
	// fmt.Printf("%#v\n", nums)

	//数组切片赋值
	// 	var arrays [10]int = [10]int{1, 2, 3, 4, 5}
	// 	nums = arrays[1:10]
	// 	fmt.Println(nums)
	// }
	//make函数 申请一个数组，初始化到切片，
	//nums = make([]int, 3)
	//	fmt.Printf("%#v %d %d", nums, len(nums), cap(nums))

	//nums = make([]int, 3, 5)
	//nums = append(nums, 2)
	//fmt.Printf("%#v %d %d", nums, len(nums), cap(nums))

	//for i := 0; i < len(nums); i++ {
	//	fmt.Println(i, nums[i])
	//}

	//for index, value := range nums {
	//		fmt.Println(index, value)
	//}

	//切片操做
	nums = make([]int, 3, 10)
	n := nums[1:3:10]
	fmt.Printf("%T %#v %d\n", n, len(n), cap(n))

	n = nums[1:4]
	fmt.Printf("%T %#v %d\n", n, len(n), cap(n))

	num1 := []int{1, 2, 3}
	num2 := []int{10, 20, 30, 40}

	copy(num1, num2) //将num2复制到num1 可以实现删除的机制
	fmt.Println(num1, num2)

	num3 := []int{1, 2, 3, 4, 5}
	fmt.Println(num3[1:])           //删除第一个元素
	fmt.Println(num3[:len(num3)-1]) //删除最后一个元素

	//删除中间某个元素
	copy(num3[2:], num3[3:])
	fmt.Println(num3[:len(num3)-1])

}
