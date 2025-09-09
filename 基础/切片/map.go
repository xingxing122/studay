package main

import (
	"fmt"
)

func main() {

	userInfo := map[string]int{
		"张三": 24,
		"李四": 28,
	}
	fmt.Println(userInfo)

	//2. 判断一个key是否在map 中

	val, ok := userInfo["zhangsan"]
	//val 返回查询可以的值   ok 返回的是bool 值，true表示存在这个key  flase 不存在这个key
	fmt.Println(val, ok)

	if _, ok := userInfo["zhangsan"]; ok {
		fmt.Println(userInfo["zhangsan"])
	} else {
		fmt.Println("map中不存在这个key")
	}
	//删除map中的一个数据

	fmt.Println(userInfo)

	delete(userInfo, "李四")

}
