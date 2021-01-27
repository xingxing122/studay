package main

import "fmt"

// 定义hello，无参，无返回值
func  sysHello(){
	fmt.Println("hello")
}

//有参数，无返回值
func  sayHi(name string){
	fmt.Println("Hi",name)
}

//有返回值 add
func  add(a int,b int)int{
	return  a  + b  //return  关键字用来向函数调用返回结果
}


func  main(){
	// 调用函数,函数名称(参数[实参])
	sysHello()
	sayHi("xingxing")
	c := add(4,3 )
    fmt.Println(c)
}
