package main

import "fmt"

func main() {

	fmt.Println("开始")
	defer fmt.Println(1)
	defer fmt.Println(3)

}
