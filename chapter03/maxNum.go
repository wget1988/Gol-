package main

import "fmt"

func main() {
	//获取最大值
	// nums := []int{1, 2, 3, 4, 5, 7, 88, 96, 22}

	// maxNumber := nums[0]

	// for i, v := range nums {
	// 	if v > maxNumber {
	// 		maxNumber = v
	// 	}
	// 	fmt.Println(i, ":", maxNumber)
	// }
	// fmt.Println(maxNumber)

	//获取第二大的值

	nums := []int{1, 2, 3, 4, 5, 7, 88, 96, 22}

	maxNumber := nums[0]
	secondNumber := nums[0]

	for i, v := range nums {
		if v > maxNumber {
			secondNumber = maxNumber
			maxNumber = v
		} else if v > secondNumber {
			secondNumber = v
		}
		fmt.Println(i, ":", maxNumber, secondNumber)
	}
	fmt.Println(maxNumber, secondNumber)

}
