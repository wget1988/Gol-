package main

import "fmt"

func main() {

	queue := []int{}

	queue = append(queue, 1)
	queue = append(queue, 2)
	queue = append(queue, 3)
	queue = append(queue, 4)
	queue = append(queue, 5)

	fmt.Println(queue)

	queue = queue[1:]
	fmt.Println(queue)
}
