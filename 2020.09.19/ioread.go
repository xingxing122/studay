package main

import "fmt"

func  main(){
	var  name string

	fmt.Print("请输入名字:")

	fmt.Scan(&name)

	fmt.Println("你输入的内容是:",name)

	var   age int
	fmt.Print("请输入你的年龄：")
	fmt.Scan(&age)
	fmt.Println("你输入的年龄是:",age)


}