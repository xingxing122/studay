package main

import (
	"fmt"
)

func main() {

	fmt.Println("老婆的想法：")
	fmt.Println("10个包子")

	var   text  string
	fmt.Print("有卖西瓜的吗:")
	fmt.Scan(&text)
	switch text {
	case "y":
		 fmt.Println("买一个西瓜")
	case  "n":
		   fmt.Println("10个包子")
	default:
		fmt.Println("10个包子")

	}

}

