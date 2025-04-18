package main

import (
	"fmt"
)

// 求两个数的和
func sumFN(x int, y int) int {
	sum := x + y
	return sum
}

// 求两个数的差
func subFn(x int, y int) int {
	sub := x - y
	return sub

}

func sumFn1(x ...int) int {

	/*	fmt.Printf("%v--%T", x, x)*/
	sum := 0
	for _, v := range x {

		sum += v
	}
	return sum
}

func calc(x, y int) (int, int) {
	sum := x + y
	sub := x - y
	return sum, sub

}

func main() {
	sum1 := sumFN(12, 3)
	fmt.Println(sum1)

	a := 20
	b := 2
	sub1 := subFn(a, b)
	fmt.Println(sub1)

	sum3 := sumFn1(12, 34, 45, 46)
	fmt.Println(sum3)

	a1, b1 := calc(100, 200)
	fmt.Println(a1, b1)

}
