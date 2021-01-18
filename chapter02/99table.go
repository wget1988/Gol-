package main

import "fmt"

func main() {

	line := 3
	for i := 1; i <= line; i++ {
		for i := 1; i <= line; i++ {
			fmt.Printf("%d * %d = %2d\t", i, line, i*line)
		}
		fmt.Println()
	}
}
