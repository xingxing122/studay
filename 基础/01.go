package main

import "fmt"

func main() {
	fmt.Print("我不会换行\n")
	fmt.Println("我是一个打印换行的方法")
	fmt.Printf("name: %s --- age: %d", "zhangsan", 28)
}

//在main函数执行前自动被调用，常用来做初始化

var log string

func init() {
	fmt.Println("init")
	log = "log对象"
}
