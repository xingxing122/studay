package main

import (
	"fmt"
)

func sumFn(x int, y int) int {
	sum := x + y
	return sum
}

func subFn(x int, y int) int {
	sub := x - y
	return sub
}

func calc(x, y int) (int, int) {
	a := x + y
	b := x - y
	return a, b
}

func main() {

	a, b := calc(100, 2)

	fmt.Println(a, b)

}
