package main

import "fmt"

func main() {

	//[]中为key类型
	var scores map[string]int //nil映射

	//定义完映射后必须初始化后再进行操作

	// fmt.Printf("%T %#v\n", scores, scores)

	// fmt.Println(scores == nil)

	scores = map[string]int{"sql": 8, "hyy": 10}
	fmt.Println(scores)

	// scores = make(map[string]int)
	// //scores = map[string]int{}
	// fmt.Println(scores)

	//增删改查
	//key不存在输出对应0值

	fmt.Println(scores["sd"])

	//判断key是否存在
	//_, ok := scores["sd"]

	if v, ok := scores["sd"]; ok {
		fmt.Println(v)
	} else {
		fmt.Println("key not found")
	}

	//修改时候，如果key存在，则修改，如果不存在，则添加

	scores["sql"] = 999
	fmt.Println(scores)

	//删除

	delete(scores, "sd")
	fmt.Println(scores)

	scores["sd"] = 100
	fmt.Println(scores)

	//遍历
	for k, v := range scores {
		fmt.Println(k, v)
	}

	//key 至少可以判断 == ！= 运算操作 bool，整数，字符串，数组
	//value 任意类型

	var myMap map[string]map[string]string
	//myMap = map[string]map[string]string{}

	myMap = map[string]map[string]string{"sql": {"地方": "兴平", "联系方式": "18700017820"}, "hyy": {"地方": "宝鸡", "联系方式": "13108049400"}}
	fmt.Println(myMap)
	fmt.Printf("%T\n %#v\n", myMap, myMap)

	fmt.Println(myMap["sql"])

	delete(myMap, "联系方式")

	fmt.Println(myMap)

}
