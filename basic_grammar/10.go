package main

import "fmt"

func main() {

	var a = 6
	var b = 3
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)

	fmt.Println(-10 % 3)

	fmt.Println(10 % 3)
	//a++
	//b--

	var a1 = 9
	var a2 = 8
	if a1 == a2 {
		fmt.Println("a1=a2")
	} else {
		fmt.Println("a1!=a2")
	}

	flag := false
	fmt.Println(!flag)

}
