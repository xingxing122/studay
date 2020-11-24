package main

import "fmt"

func main() {
	fmt.Println("老公的想法:")
	//判断，是否有卖西瓜的(控制台输入)
	//有  买1个包子
	// 否则买十个包子
	var test string
	fmt.Print("有卖西瓜的吗")
	fmt.Scan(&test)

	if test == "y" {
		fmt.Println("1个包子")
	} else {
	fmt.Println("10个包子")
  }

  }

