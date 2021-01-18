package main

import "fmt"

func main() {

	var marrays [3][2]int

	fmt.Println(marrays)
	fmt.Println(marrays[0])
	fmt.Println(marrays[0][0])
	marrays[0] = [2]int{1, 2}
	fmt.Println(marrays)

}
