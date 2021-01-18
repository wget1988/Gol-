package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for {

		//设置随机种子
		rand.Seed(time.Now().Unix())

		//生成随机数
		guess := rand.Intn(100)

		const maxGuessNumber = 5

		var isOk bool
		var i int

		for i = 0; i < maxGuessNumber; i++ {

			var input int
			fmt.Print("请输入猜的数值")
			fmt.Scan(&input)

			if guess == input {
				fmt.Printf("right %d次就猜对了\n", i+1)
				isOk = true
				break
			} else if input > guess {
				fmt.Printf("数字太大，还有%d次机会", maxGuessNumber-i-1)
			} else {
				fmt.Printf("猜得太小了，还有%d次机会", maxGuessNumber-i-1)
			}
		}
		if !isOk {
			fmt.Printf("太笨了")
		}
		var text string
		fmt.Print("是否继续y/n")
		fmt.Scan(&text)
		if text != "y" {
			break
		}
	}

}
