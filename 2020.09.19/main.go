package main

import "fmt"
// 全局变量
var  name  string
//若未设置值则使用零值进行初始化

func main() {
//局部变量
var (
	age int  = 30
	weight  int = 40
)



height  := 168
fmt.Println(name,age,weight,height)


}
