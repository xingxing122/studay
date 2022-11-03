package main

import "fmt"

func main() {
	fmt.Println("老公的想法")

	var text string
	fmt.Print("有卖西瓜的吗")
	fmt.Scan(&text)

	if text == "y" {
		fmt.Println("1个包子")

	} else {
		fmt.Println("10个包子")
	}

}
