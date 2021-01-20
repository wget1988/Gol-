package main

import "fmt"

func main() {
	users := map[string]int{"sql": 100, "hyy": 88}

	keySlice := make([]string, len(users))

	//	keySlice[0] = "" //切片索引赋值
	valueSlice := []int{}

	i := 0
	for k, v := range users {
		keySlice[i] = k
		i++

		valueSlice = append(valueSlice, v)
	}
	fmt.Println(keySlice, valueSlice)

	for k, _ := range users {
		fmt.Println(k)
	}

	for _, v := range users {
		fmt.Println(v)
	}
	//如果k v 中只写一个，则默认为k
}
