package main

import "fmt"

func cal(num1 int, num2 int) int {

	var sum int = 0
	sum += num1
	sum += num2
	return sum
}

func main() {
	// 调用 cal 并将返回值赋给 sum
	sum := cal(10, 20)
	fmt.Println(sum) // 打印结果
}