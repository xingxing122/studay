package main

import "fmt"

func main() {

	var score int = 93
	if score >= 90 {
		fmt.Println("你的成绩为A级别")
	}
	if score >= 80 {
		fmt.Println("你的成绩为B级别")
	}
	if score >= 70 {
		fmt.Println("你的成绩为C级别")
	}
}
