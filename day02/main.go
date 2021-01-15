package main

import "fmt"

func main() {
	// var i = 100
	// fmt.Printf("%d\n", i)

	// i2 := 077
	// fmt.Printf("%d\n", i2)

	// i3 := 0x1234567
	// fmt.Printf("%d\n", i3)

	// // 字符串拼接
	// name := "study"
	// age := "16"
	// ss := name + age
	// fmt.Printf(ss)
	// ss1 := fmt.Sprintf("\n%s%s", name, age)
	// fmt.Println(ss1)

	//分割
	//ret := strings.Split(s3, "\\")

	//包含
	//fmt.Println(strings.Contains(ss, "study"))

	//前缀
	//fmt.Println(strings.HasPrefix(ss, "study"))

	//后缀
	//fmt.Println(strings.HasSuffix(ss, "study"))

	//拼接
	//fmt.Println(strings.Join(ret, "+"))

	str := "helloworld"

	for _, c := range str {
		fmt.Printf("%c\n", c)
	}

	// go get -d github.com/ramya-rao-a/go-outline
	// go get -d github.com/acroca/go-symbols
	// go get -d github.com/mdempsky/gocode
	// go get -d github.com/rogpeppe/godef
	// go get -d github.com/zmb3/gogetdoc
	// go get -d github.com/fatih.gomodifytags
	// go get -d sourcegraph.com/sqs/goreturns
	// go get -d github.com/cweill/gotests/...
	// go get -d github.com/josharian/impl
	// go get -d github.com/haya14busa/goplay/cmd/goplay
	// go get -d github.com/uudashr/gopkgs/cmd/gopkgs
	// go get -d  github,com/davidrgenni/reftools/cmd/filestruct
	// go get -d github.com/alecthomas/gometalinter

	//安装，将上述命令去掉-d
	// gometalinter --install // 安装扩展

}
