package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "hello world"
	s_zh := "中国"
	fmt.Println(len(s))
	fmt.Println(len(s_zh))

	str := "11+12+13"
	arr1 := strings.Split(str, "+")
	fmt.Println(arr1)

}
