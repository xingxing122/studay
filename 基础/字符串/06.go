package main

import (
	"fmt"
	"strconv"
)

//字符串转换

func main() {
	//1. 将int 转换为字符串
	num := 100

	fmt.Printf("%T %d \n", num, num)

	strNum := strconv.Itoa(num)
	fmt.Printf("%T %v \n", strNum, strNum)

	//将字符串转为int
	intNum, err := strconv.Atoi(strNum)
	if err != nil {
		fmt.Println(err)

	}
	fmt.Printf("%T: %v", intNum, intNum)
}
