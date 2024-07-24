package main

import "fmt"

func main() {

	var y string

	fmt.Print("有卖西瓜的吗:")
	fmt.Scan(&y)

	switch y {
	case "yes", "y", "Y":
		fmt.Println("买一个包子")
	case "no", "n", "N":
		fmt.Println("买十个包子")
	default:
		fmt.Println("输入错误")
	}

}
