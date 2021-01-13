package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
    var  name string  = "xingxing"
    var  desc string ="i love china"
    // 字面量
    // 零值
    // 操作
    // 连接
    // 关系运算
    // 索引，字节
    // 切片 => 字符串，字节

    fmt.Println(name,desc)
    fmt.Println(desc[1:5])

    for i, v := range  desc {
    	fmt.Println(i,v)
	}

	var  txt = "我爱中国"
	for i , v := range txt {
		fmt.Printf("%d,%q\n", i,v )
	}
	fmt.Println(utf8.RuneCountInString(txt))  // 打印字节




}
