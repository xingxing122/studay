package main

import "fmt"

func main() {
	fmt.Println("老婆的想法:")
	fmt.Println("10个包子")
	// 判断，是否有卖西瓜的
	// 有  买一个西瓜
	var   test   string
	fmt.Print("有卖西瓜的吗:")
	fmt.Scan(&test)
	
	if  test  == "y" {
		fmt.Println("1个西瓜")
	}



}
