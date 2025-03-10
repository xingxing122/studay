package main

import (
	"fmt"
	"strconv"
)

func main() {

	//通过strconv  把其他类型转换为string 类型
	var i int = 20
	str1 := strconv.FormatInt(int64(i), 10)
	fmt.Printf("值: %v 类型： %T\n", str1, str1)

	//
	var f float32 = 20.23
	var str2 = strconv.FormatFloat(float64(f), 'f', 4, 64)
	fmt.Printf("值: %v 类型：%T\n", str2, str2)

}
