package main

import "fmt"

func main() {

	var n1 int = 100
	fmt.Println(n1)

	var n2 float32 = float32(n1)
	fmt.Println(n2)
	//注意，N1 的类型还是int 类型，只是将N1的值100 转为了float32而已，N1 还是int 类型

	fmt.Printf("%T", n1)

	var n3 int64 = 88888
	var n4 int8 = int8(n3)
	fmt.Println(n4)

	var n5 int32 = 12
	var n6 int64 = int64(n5) + 30
	fmt.Println(n5)
	fmt.Println(n6)
	
}
