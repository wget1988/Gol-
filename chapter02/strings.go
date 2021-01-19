package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Compare("abc", "asd"))
	fmt.Println(strings.Contains("abc", "ad"))
	fmt.Println(strings.ContainsAny("abc", "ae"))
	fmt.Println(strings.Count("asdssdad", "asd"))
	fmt.Println(strings.Fields("abc de sad \neeeeee\raaaa\fxx\ad"))

	fmt.Println(strings.Split("asdaidid;adq;adsadwqe;wqewqedasdasdl", ";"))
	fmt.Println(strings.Join([]string{"asd", "dad", "sdadad"}, ";"))

	fmt.Println(strings.Repeat("adsdasada", 2))
	fmt.Println(strings.Replace("abcsaiddadadadad", "ad", "xxx", 1))
	fmt.Println(strings.Replace("abcsaiddadadadad", "ad", "xxx", -1)) //替换所有
}
