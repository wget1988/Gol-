package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	MAX_AUTH = 3

	PASSWORD = "shiqinlong"
)

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

func inputString(prompt string) string {
	var input string
	fmt.Print(prompt)
	fmt.Scan(&input)
	return strings.TrimSpace(input)
}

//查询用户信息

func printUser(pk int, users map[string]string) {

	fmt.Println("ID", pk)
	fmt.Println("姓名：", users["name"])
	fmt.Println("年龄：", users["age"])
	fmt.Println("电话：", users["tel"])
	fmt.Println("地址：", users["addr"])

}
func query(users map[int]map[string]string) {

	q := inputString("请输入查询信息")
	title := fmt.Sprintf("%5s|%20s|%5s|%20s|%50s", "ID", "NAME", "AGE", "TEL", "ADDR")
	fmt.Println(title)
	fmt.Println(strings.Repeat("-", len(title)))
	for k, user := range users {
		if strings.Contains(user["name"], q) || strings.Contains(user["add"], q) || strings.Contains(user["tel"], q) || strings.Contains(user["age"], q) {
			printUser(k, user)
			fmt.Println("--------------------")
			fmt.Println()
		}
	}
}

func getId(users map[int]map[string]string) int {
	var id int
	for k := range users {
		if id < k {
			id = k
		}
	}
	return id + 1
}

//添加用户信息
func add(users map[int]map[string]string) {
	id := getId(users)

	user := inputUser()
	users[id] = user
	fmt.Println("添加成功")
}

//输入用户信息
func inputUser() map[string]string {
	user := map[string]string{}
	user["name"] = inputString("请输入名字：")
	user["age"] = inputString("请输入年龄：")
	user["tel"] = inputString("请输入电话：")
	user["addr"] = inputString("请输入地址：")

	return user
}

//修改
func modify(users map[int]map[string]string) {
	idString := inputString("请输入修改用户ID：")
	if id, err := strconv.Atoi(idString); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将修改的用户信息")
			printUser(id, user)
			input := inputString("确定修改Y/N")
			if input == "y" || input == "Y" {
				user := inputUser()
				users[id] = user
				fmt.Println("修改完成")
			}
		} else {
			fmt.Println("ID不存在")
		}
	} else {
		fmt.Println("输入ID不正确")
	}
}

//删除

func deleteUser(users map[int]map[string]string) {
	idString := inputString("请输入删除用户ID：")
	if id, err := strconv.Atoi(idString); err == nil {
		if user, ok := users[id]; ok {
			fmt.Println("将删除的用户信息")
			printUser(id, user)
			input := inputString("确定删除Y/N")
			if input == "y" || input == "Y" {
				delete(users, id)
				fmt.Println("删除完成")
			}
		} else {
			fmt.Println("ID不存在")
		}
	} else {
		fmt.Println("输入ID不正确")
	}
}

func main() {
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
	//title := fmt.Sprintf("%5s|%20s|%5s|%20s|%50s", "ID", "NAME", "AGE", "TEL", "ADDR")

	users := map[int]map[string]string{}

	callbacks := map[string]func(map[int]map[string]string){
		"1": add,
		"2": modify,
		"3": deleteUser,
		"4": query,
	}
	//var id int
	// users := map[int][5]string
	// users := map[int][int]string
	// users := []map[string]string{}
	// users := [][]string
	// users := [][5]string

	fmt.Println("欢迎登录")
	//END:
	for {
		//	fmt.Println("请输入指令：（create/delete/replace/show/exit）")
		//fmt.Scan(&string)

		fmt.Println(menu)
		op := inputString("请输入指令：")
		callback, ok := callbacks[op]

		if ok {
			callback(users)
		} else if op == "5" {
			break
		} else {
			fmt.Println("指令错误")
		}
		// switch op {
		// case "1":
		// 	add(users)
		// case "2":
		// 	modify(users)
		// case "3":
		// 	deleteUser(users)
		// case "4":
		// 	query(users)
		// case "5":
		// 	//return
		// 	break END
		// default:
		// 	fmt.Println("指令错误")
		// }
	}
}
