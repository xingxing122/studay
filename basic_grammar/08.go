package main

import "fmt"

func main() {
	//类型转换，整形与浮点型
	var a int8 = 20
	var b int16 = 40

	fmt.Println(int16(a) + b)

	//浮点型
	var c float32 = 30
	var d float64 = 40

	fmt.Println(float64(c) + d)

	//浮点型与整型
	var h float32 = 40.24
	var g int = 50

	fmt.Println(h + float32(g))
	//从低位转换为高位

}
