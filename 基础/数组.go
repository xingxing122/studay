package main

import "fmt"

//数组

func main() {

	var a [5]int

	fmt.Println(a)
	a[0] = 100
	// 数组是有长度限制的

	fmt.Println(a)

	var b = [...]string{"bj", "sh", "gz"}
	fmt.Println(b)

	//普通方法通过索引遍历数组
	for i := 0; i < len(a); i++ {
		fmt.Println(i, a[i])
	}
	//for  range 方法循环数组
	//index 是索引， value 获取的值， := range 是go 中自己实现的方法
	for index, value := range a {
		fmt.Println(index, value)
	}

}
