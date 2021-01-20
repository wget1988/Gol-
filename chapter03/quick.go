package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{1, 3, 44, 2, 5, 7, 9}

	sort.Ints(nums)

	fmt.Println(nums)

	fmt.Println(sort.SearchInts(nums, 5))
	//二分查找 在有序的数组中查找元素
	num := 5
	if nums[sort.SearchInts(nums, num)] == num {
		fmt.Println("exist")
	}

}
