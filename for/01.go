package main

import "fmt"

func main() {
	var text string
	fmt.Print("有卖西瓜的吗:")
	fmt.Scan(&text)

	if text == "y" {
		fmt.Println("1个西瓜")
	}

}
