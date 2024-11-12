package main

import "fmt"

func main() {

	var score int = 187

	switch score / 2 {
	case 2:
		fmt.Println("你的等级为A")
	case 1:
		fmt.Println("你的等级为B")
	default:
		fmt.Println("你的成绩有误")
	}

}
