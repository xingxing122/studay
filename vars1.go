package main

import "fmt"

func main() {
	//如果有卖西瓜的，卖一个包子，

	var y string

	fmt.Print("有卖西瓜的吗:")
	fmt.Scan(&y)

	if y == "yes" {
		fmt.Println("买一个包子")
	} else {

		fmt.Println("买10个包子")
	}

}
