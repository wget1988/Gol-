package main

import (
	"fmt"
	"sort"
)

func main() {

	nums := []int{1, 23, 5, 6, 645, 234}
	sort.Ints(nums)
	fmt.Println(nums)

	names := []string{"test", "work", "awk"}
	sort.Strings(names)
	fmt.Println(names)
}
