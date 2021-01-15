package main

import "fmt"

//var version string = "ssh"

func main() {

	//声明单个变量
	//var me string
	//me = "shiqinlong"
	//fmt.Println(me)
	//fmt.Println(version)

	//声明多个变量
	/*
		var (
			//不同类型
			user string
			name string
			age  int
		)
	*/
	/*
		//如过不写变量类型，必须赋值，go会通过默认的值去推导其类型
		var (
			student   = "sss"
			school    = "xupt"
			classroom = "11"
		)

		fmt.Println(student, classroom, school)
	*/
	///下面是简短声明，只能在函数内部使用
	//isBoy := true
	//fmt.Println(isB
	//	}
	//const NAME= "shiqinlong "
	// 	const AGE int = 24
	// 	const (
	// 		STU = "shiqinlong"
	// 		ID  = 1001
	// 	)
	// 	fmt.Println(AGE, STU, ID)
	// 	//若后面省略类型和值，则默认与上一个常量值一致，常应用于枚举类型
	// 	const (
	// 		C1 int = 1
	// 		C2
	// 		C3
	// 	)
	// }
	outer := 1
	{
		inner := 2
		fmt.Println(outer)
		fmt.Println(inner)
		{
			//outer := 4
			inner2 := 3
			fmt.Println(outer, inner, inner2)
		}
		inner2 := 2
		fmt.Println(inner2)
	}

}
