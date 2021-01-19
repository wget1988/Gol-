package main

import "fmt"

func main() {
	heights := []int{10, 6, 7, 9, 5}

	//把最高的移动到最后
	//从第一个人开始，两两来是比较，如果前面人高，两人交换位置交换
	//第一次：10 6 => 6 10 7 9 5
	//第二次：10 7 => 6 7 10 9 5
	//第三次： 10 9 => 6 7 9 10 5
	//第四次： 10 5 =》 6 7 9 5 10
	//n 个人比较 n-1

	//第二轮： ......
	for j := 0; j < len(heights)-1; j++ {
		fmt.Println("第", j, "轮")
		for i := 0; i < len(heights)-1; i++ {
			if heights[i] > heights[i+1] {
				fmt.Println("交换：", heights[i], heights[i+1])
				tmp := heights[i]
				heights[i] = heights[i+1]
				heights[i+1] = tmp
			}
			fmt.Println(i, "交换完毕", heights)
		}
		fmt.Println("第", j, "轮，结果", heights)
	}
}
