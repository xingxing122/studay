package main

import "fmt"

//函数递归调用

func fn1(n int) {
	if n > 0 {
		fmt.Println(n)
		n--
		fn1(n)
	}
}

func fn2(n int) int {
	if n > 1 {
		return n + fn2(n-1)
	} else {
		return 1
	}

}

func main() {
	var sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)
	fn1(5)

	fmt.Println(fn2(100))
}
