package main

import "fmt"

func main() {

	// 01
	//age := 30
	//if age > 20 {
	//	fmt.Println("成年人")
	//}

	//02

	//if a := 34; a > 30 {
	//	fmt.Println("成年人", a)
	//}
	//03
	// var score = 75
	//	if score >= 90 {
	//		fmt.Println("A")
	//	} else if score > 75 {
	//		fmt.Println("B")
	//	} else {
	//		fmt.Println("C")
	//	}

	//求两个数的最大值

	var d = 34
	var f = 24
	var max int
	if d > f {
		max = d
	} else {
		max = f
	}
	fmt.Println("d和f 的最大值是多少", max)

}
