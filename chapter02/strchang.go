package main

import (
	"fmt"
	"strconv"
)

func main() {

	//v, err := strconv.ParseBool("true") //返回两个值，bool值和是否有错误
	//fmt.Println(v, err)
	//字符串转换为bool
	if v, err := strconv.ParseBool("true"); err == nil {
		fmt.Println(v)
	} else {

		fmt.Println(err)
	}

	//int
	if v, err := strconv.Atoi("1023"); err == nil {
		fmt.Println(v)
	} else {
		fmt.Println(err)
	}

	if v, err := strconv.ParseInt("0x64", 16, 64); err == nil {
		fmt.Println(v)
	}
	//float
	if v, err := strconv.ParseFloat("1.1", 64); err == nil {
		fmt.Println(v)
	}

	sd := fmt.Sprintf("%d", 12) //int->字符串
	fmt.Println(sd)

	sf := fmt.Sprintf("%.3f", 12.1)
	fmt.Println(sf)

	fmt.Println(strconv.FormatBool(false)) //Format将其他类型转换为字符串

}
