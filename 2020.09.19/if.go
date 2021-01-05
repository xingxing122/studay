package main

import "fmt"

func main() {
	fmt.Println("老婆的想法")
	fmt.Println("10个包子")


	var  test  string
	fmt.Println("有卖西瓜的吗")
	fmt.Scan(&test)

	if  test  == "y" {
		fmt.Println("买一个西瓜")
	} else {
		fmt.Println("10个包子")
	}
	
}
