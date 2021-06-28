package main

import "fmt"

func main() {

	fmt.Println("老婆的想法:")
	fmt.Println("10个包子")


	var  test string
	fmt.Print("有卖西瓜的吗:")
    fmt.Scan(&test)

	if test  == "y" {
		fmt.Println("1个包子")
	} else {
		fmt.Println("买10个包子")
	}



}
