package main

import "fmt"

func main() {
	//声明引用类型，必须要使用make 或new 方法申请内存空间
	//make 方法申请内存空间
	// 第一个[]int 指定数据类型是一个int 切片
	// 3 10 切片默认长度是3，默认容量是10
	a := make([]int, 3, 10)
	fmt.Println(a)

	//通过new 申请空间
	b := new([]int)

	fmt.Printf("%T: %T", a, b)

}
