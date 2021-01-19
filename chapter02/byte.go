package main

import "fmt"

func main() {

	var byteS []byte = []byte{'a', 'b', 'c'}
	fmt.Printf("%T,%#v\n", byteS, byteS)

	s := string(byteS) //字节转换字符串

	fmt.Printf("%T %#v", s, s)

	bs := []byte(s)
	fmt.Printf("%T %#v", bs, bs)
}
	