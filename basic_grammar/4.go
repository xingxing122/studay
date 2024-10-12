package main

import (
	"fmt"
	"strconv"
)

func main() {
	var n1 int = 18
	var s1 string = strconv.FormatInt(int64(n1), 10) //参数，第一个参数必须转为int64 类型，第二个参数指定字面值的进制形式转为十进制

	fmt.Printf("s1 对应的类型是：%T, s1=%q\n", s1, s1)

	var s2 string = "true"
	var b bool
	b, _ = strconv.ParseBool(s2)
	fmt.Println(b)

}
