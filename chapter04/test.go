package main

import "fmt"

func main() {

	array := [3]string{"A", "B", "C"}

	slice := []string{"A", "B", "C"}

	arrayA := array
	sliceA := slice

	fmt.Println(array, arrayA)
	fmt.Println(slice, sliceA)

	arrayA[0] = "X"
	sliceA[0] = "Y"
	fmt.Println(array, arrayA)
	fmt.Println(slice, sliceA)

	m := map[string]string{}

	mA := m
	mA["sql"] = "xian"
	fmt.Println(m, mA)

	point := &array[0]

	*point = "add"

	fmt.Println(array, *point)
} 
