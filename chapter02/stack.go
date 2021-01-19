package main

import "fmt"

func main() {

	queue := []int{}
	// for _,value:=1 range queue {
	// 	queue = append(queue,value)
	// 	queue +
	// }
	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	queue = append(queue, 4)
	queue = append(queue, 5)

	fmt.Println(queue)

	queue = queue[:len(queue)-1]
	fmt.Println(queue)

	point := [][]int{}

	point1 := make([][]int, 0)

	fmt.Printf("%T\n", point1)

	point = append(point1, []int(1, 2, 3, 4))
	fmt.Println(point)
	fmt.Println(point[0])
	fmt.Println(point1[0][1])
}
