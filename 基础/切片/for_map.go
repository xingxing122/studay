package main

import (
	"fmt"
)

func main() {

	userInfo := map[string]int{
		"张三":     24,
		"wangwu": 32,
	}
	//1.使用for range 只遍历key

	for k := range userInfo {
		fmt.Println(k)
		if val, ok := userInfo[k]; ok {

			fmt.Println(k, val)
		}
	}
	//2.遍历key 和value
	for key1, value1 := range userInfo {
		fmt.Println(key1, value1)

	}

}
