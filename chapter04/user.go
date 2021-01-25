package main

import (
	"fmt"
	"strings"
)

const (
	MAX_AUTH = 3

	PASSWORD = "shiqinlong"
)

//添加用户
func addUser(key int, users map[string]map[string]string) {
	var (
		id string = fmt.Sprintf("%d", key)
		//	id   string = strconv.Itoa(key)
		name string
		age  string
		tel  string
		addr string
	)

	fmt.Print("请输入姓名：")
	fmt.Scan(&name)

	fmt.Print("请输入年龄：")
	fmt.Scan(&age)

	fmt.Print("请输入电话：")
	fmt.Scan(&tel)

	fmt.Print("请输入地址：")
	fmt.Scan(&addr)

	users[id] = map[string]string{
		"id":   id,
		"name": name,
		"age":  age,
		"tel":  tel,
		"addr": addr,
	}
}

//查询用户信息
func query(users map[string]map[string]string) {

	//q为空，查全部
	//非空，必须名称，电话，住址中包含q内容显示
	var q string

	fmt.Print("请输入查询信息")
	fmt.Scan(&q)
	title := fmt.Sprintf("%5s|%20s|%5s|%20s|%50s", "ID", "NAME", "AGE", "TEL", "ADDR")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))

	for _, user := range users {
		if q == "" || strings.Contains(user["name"], q) || strings.Contains(user["add"], q) || strings.Contains(user["tel"], q) || strings.Contains(user["age"], q) {
			fmt.Printf("%5s|%20s|%5s|%20s|%50s", user["id"], user["name"], user["age"], user["tel"], user["addr"])
			fmt.Println()
		}

	}
}

//用户认证 ，从命令行输入密码，验证通过返回值告知用户成功或者失败

func auth() bool {

	var password string
	for i := 0; i < MAX_AUTH; i++ {
		fmt.Println("请输入密码：")
		fmt.Scan(&password)
		if PASSWORD == password {
			return true
		} else {
			fmt.Println("密码错误")
		}
	}
	return false
}

func main() {
	//存储用户信息
	users := make(map[string]map[string]string)

	id := 0

	if !auth() {
		fmt.Printf("密码%d次错误，系统退出\n", MAX_AUTH)
		return
	}

	menu := `
	********************
	1. 新建
	2. 修改
	3. 删除
	4. 查询
	5. 退出
	********************
	`
	fmt.Println("欢迎登录")
END:
	for {
		var op string
		//	fmt.Println("请输入指令：（create/delete/replace/show/exit）")
		//fmt.Scan(&string)

		fmt.Println(menu)
		fmt.Println("请输入指令：")
		fmt.Scan(&op)
		switch {
		case op == "1":
			id++
			addUser(id, users)
		case op == "2":
		case op == "3":
		case op == "4":
			query(users)
		case op == "5":
			//return
			break END
		default:
			fmt.Println("指令错误")
		}
	}
}
