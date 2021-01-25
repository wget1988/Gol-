package main

import "fmt"

func test() (err error) {
	defer func() {
		if e := recover(); e != nil {
			//fmt.Println(err)
			err = fmt.Errorf("%v", e)
		}
	}()

	panic("error")
	return
}

func main() {
	// defer func() {
	// 	if err := recover(); err != nil {
	// 		fmt.Println(err)
	// 	}
	// }()

	// fmt.Println("main start")
	// panic("error")

	// fmt.Println("over")
	err := test()
	fmt.Println(err)
}
